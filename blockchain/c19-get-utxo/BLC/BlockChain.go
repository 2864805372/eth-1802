package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
	"math/big"
	"os"
	"strconv"
)

// 区块链管理相关
// 持久化数据库名称，就是一个数据库文件
const dbName = "block.db"

// 表(桶)名称
const blockTableName = "blocks"

// 区块链基本结构 cursor
type BlockChain struct {
	Tip		[]byte		// 保存最新区块的哈希值
	DB 		*bolt.DB	// 数据库对象
}

// 判断数据库文件存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		// 文件不存在
		return false
	}
	return true
}

// 初始化区块链
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {
	if dbExists() {
		fmt.Println("创世区块已存在...")
		os.Exit(1)
	}
	var blockHash []byte
	// 创建或打开一个数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("open the db failed! %v\n", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil == b { // 没找到表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if nil != err {
				log.Panicf("create the bucket [%s] failed! %v\n", blockTableName, err)
			}
		}

		if nil != b {
			// 生成交易
			txCoinbase := NewCoinbaseTransaction(address)
			// 在bucket存在的情况下，创建创世区块
			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinbase})
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if nil != err {
				log.Panicf("insert the genesis block to boltdb failed1 %v\n", err)
			}

			// 存储最新区块的哈希
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if nil != err {
				log.Panicf("put the genesis block hash to boltdb failed1 %v\n", err)
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	if nil != err {
		log.Panicf("update the data of genesis block failed! %v\n", err)
	}

	return &BlockChain{blockHash, db}
}

// 添加新的区块到区块链中
func (bc *BlockChain) AddBlock(txs []*Transaction)  {
	//newBlock := NewBlock(height,prevBlockHash, data)
	//bc.Blocks = append(bc.Blocks, newBlock)
	// 更新数据
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 1. 获取数据表
		b := tx.Bucket([]byte(blockTableName))
		if nil != b { // 2. 确保表确实存在
			blockBytes := b.Get(bc.Tip)
			// 3. 获取数据库中最后插入的区块
			latest_block := DeserializeBlock(blockBytes)
			// 4. 创建新区块
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, txs)
			// 5. 存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("put the new block into the db failed! %v\n", err)
			}
			// 6. 更新最新区块的哈希
			err = b.Put([]byte("l"), newBlock.Hash)
			if nil != err {
				log.Panicf("put the latest block hash into the db failed! %v\n", err)
			}
			bc.Tip = newBlock.Hash
		}

		return nil
	})

	if nil != err {
		log.Panicf("update the db when insert the new block failed! %v\n", err)
	}
}

// 遍历输出区块链所有区块的信息
func (bc *BlockChain) PrintChain()  {
	fmt.Println("区块链完整信息...")
	var curBlock *Block		// 当前区块结构
	bcit := bc.Iterator() // 创建一个迭代器对象
	for  {
		fmt.Println("---------------------------------------------------")
		curBlock = bcit.Next() // 直接迭代

		fmt.Printf("\tHeight : %d\n", curBlock.Height)
		fmt.Printf("\tTimeStamp : %d\n", curBlock.TimeStamp)
		fmt.Printf("\tPrevBlockHash : %x\n", curBlock.PrevBlockHash)
		fmt.Printf("\tHash : %x\n", curBlock.Hash)
		fmt.Printf("\tNonce : %d\n", curBlock.Nonce)
		//fmt.Printf("\tTransaction : %v\n", curBlock.Txs)
		for _, tx := range curBlock.Txs {
			fmt.Printf("\t\ttx-hash : %x\n", tx.TxHash)
			fmt.Printf("\t\t输入...\n")
			for _, vin := range tx.Vins {
				fmt.Printf("\t\t\tvin-txhash:%x\n", vin.TxHash)
				fmt.Printf("\t\t\tvin-vout:%d\n", vin.Vout)
				fmt.Printf("\t\t\tvin-scriptSig:%s\n", vin.ScriptSig)
			}
			fmt.Printf("\t\t输出...\n")
			for _, vout := range tx.Vouts {
				fmt.Printf("\t\t\tvout-value:%d\n", vout.Value)
				fmt.Printf("\t\t\tvout-ScriptPubkey:%s\n", vout.ScriptPubkey)
			}
		}
		// 判断是否已经遍历创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break // 跳出循环
		}
	}
}

// 返回blockChain 对象
func BlockchainObject() *BlockChain {
	// 读取数据库
	db, err := bolt.Open(dbName,0600, nil)
	if nil != err {
		log.Panicf("open the db of blockchain failed! %v\n", err)
	}

	// 获取最新区块哈希值
	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if nil != err {
		log.Panicf("get the object of blockchain failed! %v\n", err)
	}
	// 赋值
	return &BlockChain{tip, db}
}

// 挖矿(生成新区块)
// 通过接收指定交易，进行打包确认，最终生成新的区块
func (blockchain *BlockChain)MineNewBlock(from, to, amount []string) {
	// 交易列表
	var txs []*Transaction
	value, _ := strconv.Atoi(amount[0])
	// 生成新的交易
	tx := NewSimpleTransaction(from[0],to[0], value)
	// 追加到交易列表
	txs = append(txs, tx)
	// 生成新的区块
	var block *Block
	// 从数据库获取最新区块
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			hash := b.Get([]byte("l")) // 获取最新区块的哈希值
			blockBytes := b.Get(hash)
			block = DeserializeBlock(blockBytes)
		}
		return nil
	})
	// 生成新的区块
	block = NewBlock(block.Height + 1, block.Hash, txs)
	// 持久化新的区块
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			err := b.Put(block.Hash, block.Serialize())
			if nil != err {
				log.Panicf("update the new block to db failed! %v\n", err)
			}

			err = b.Put([]byte("l"), block.Hash)
			if nil != err {
				log.Panicf("update the latest hash to db failed! %v\n", err)
			}

			blockchain.Tip = block.Hash
		}
		return nil
	})
}

// 查找指定地址已花费输出
func (blockchain *BlockChain)SpentOutputs(address string) map[string][]int {
	bcit := blockchain.Iterator()
	// 2. 遍历每一个查找到的交易中的vout
	// 定义已花费输出列表,存储所有的已花费输出
	// key: 每个 input 所引用的交易的哈希
	// value:output 索引列表
	spentTXOutputs := make(map[string][]int)
	for  {
		block := bcit.Next()
		for _, tx := range block.Txs {
			// 先查找输入，排除coinbase交易
			if !tx.IsCoinbaseTransaction() {
				for _, in := range tx.Vins {
					// 判断能不能引用指定地址的输出
					if in.UnLockWithAddress(address) {
						// 添加到已花费输出的map里
						key := hex.EncodeToString(in.TxHash)
						spentTXOutputs[key] = append(spentTXOutputs[key],in.Vout)
					}
				}
			}
		}
		// 退出循环条件
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return spentTXOutputs
}

// 返回指定地址UTXO(未花费输出)
func (blockchain *BlockChain)UnUTXOS(address string) []*TxOutput {
	fmt.Printf("the address is %s\n", address)
	var unUTXOS []*TxOutput
	// 1. 遍历区块链，查找与address相关的所有交易
	// 获取迭代器对象
	bcit := blockchain.Iterator()
	// 2. 遍历每一个查找到的交易中的vout
	// 定义已花费输出列表,存储所有的已花费输出
	// key: 每个 input 所引用的交易的哈希
	// value:output 索引列表
	spentTXOutputs := blockchain.SpentOutputs(address)
	for _, spent := range spentTXOutputs {
		fmt.Printf("spent : %v", spent)
	}
	// 迭代
	for  {
		// 获取每一个区块信息
		block := bcit.Next()
		// 遍历每个区块中的交易
		for _, tx := range block.Txs {
			// 再查找输出
			for index, vout := range tx.Vouts {
				// 地址验证(检查输出是否属于传入的账号)
				// 判断out是否是一个未花费的输出
				if len(spentTXOutputs) != 0 {
					// txHash = 111111
					// indexArray = [0]
					for txHash, indexArray := range spentTXOutputs{
						for _, i := range indexArray {
							// tx.TxHash = 222222
							if txHash ==hex.EncodeToString(tx.TxHash) && i == index {
								continue
							} else {
								unUTXOS = append(unUTXOS, vout)
							}
						}
					}
				} else {
					unUTXOS = append(unUTXOS, vout)
				}
			}
		}
		// 退出循环条件
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	// 3. 判断vout是否被花费，另外，判断指定vout是否属于address
	return unUTXOS
}

// 查询指定地址的余额
func (blockchain *BlockChain) getBalance(address string) int64 {
	// 查找指定地址UTXO
	utxos := blockchain.UnUTXOS(address)
	var amount int64
	for _, utxo := range utxos {
		// 获取每个UTXO的VALUE，累加
		amount += utxo.Value
	}
	return amount
}