package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
	"math/big"
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

// 初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	var blockHash []byte
	// 创建或打开一个数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("open the db failed! %v\n", err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			latest_hash := b.Get([]byte("l"))
			if latest_hash != nil {
				blockHash = latest_hash
			}
		}
		return nil
	})
	if nil != err {
		log.Panicf("view the bucket failed! %v\n", err)
	}
	if blockHash != nil {
		return &BlockChain{blockHash, db}
	}
	//defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil == b { // 没找到表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if nil != err {
				log.Panicf("create the bucket [%s] failed! %v\n", blockTableName, err)
			}

			// 在bucket存在的情况下，创建创世区块
			genesisBlock := CreateGenesisBlock("the data of genesis block")
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
		} else {
			fmt.Printf("the genesis block is already exist\n")
		}
		return nil
	})
	if nil != err {
		log.Panicf("update the data of genesis block failed! %v\n", err)
	}

	return &BlockChain{blockHash, db}
}

// 添加新的区块到区块链中
func (bc *BlockChain) AddBlock(data []byte)  {
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
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, data)
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
		fmt.Printf("\tData : %x\n", curBlock.Data)
		fmt.Printf("\tNonce : %d\n", curBlock.Nonce)
		// 判断是否已经遍历创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break // 跳出循环
		}
	}
}