package BLC

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
)

// UTXO持久化相关管理

// utxo table name
const utxoTableName = "utxoTable"

// UTXOSet结构(保存指定区块链中所有UTXO)
type UTXOSet struct {
	BlockChain *BlockChain
}

// 将UTXO集合序列化为字节数组
func (txOutputs *TXOutputs) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(txOutputs); nil != err {
		log.Panicf("serialize the utxo failed! %v\n", err)
	}
	return result.Bytes()
}

// 反序列化
func DeserializeTXOutputs(txOutputsBytes []byte) *TXOutputs {
	var txOutputs TXOutputs
	decoder := gob.NewDecoder(bytes.NewReader(txOutputsBytes))
	err := decoder.Decode(&txOutputs)
	if nil != err {
		log.Panicf("deserialize the struct of txOutputs failed! %v\n", err)
	}
	return &txOutputs
}
// 重置UTXO，可以在创建区块链的时候调用
func (utxoSet *UTXOSet)ResetUTXOSet()  {
	// 在第一创建的时候更新utxo table
	// 采用覆盖的方式，如果指定表原本存在，则先删除
	err := utxoSet.BlockChain.DB.Update(func(tx *bolt.Tx) error {
		// 查找UTXO表
		b := tx.Bucket([]byte(utxoTableName))
		// 如果表原本存在，则先删除
		if nil != b {
			fmt.Println("delete bucket")
			err := tx.DeleteBucket([]byte(utxoTableName))
			if nil != err {
				panic(err)
			}
		}
		// 创建utxoTable
		c, _ := tx.CreateBucket([]byte(utxoTableName))
		if nil != c {
			// 查找所有未花费输出
			txOutputsMap := utxoSet.BlockChain.FindUTXOMap()
			fmt.Printf("txOutputsMap:%v\n",txOutputsMap )
			for keyHash, output := range txOutputsMap {
				fmt.Printf("keyHash : %s\n", keyHash)
				txHash, _ := hex.DecodeString(keyHash)
				// 存入utxo table
				err := c.Put(txHash, output.Serialize())
				if nil != err {
					log.Panicf("put the utxo into table failed! %v\n", err)
				}
			}
		}
		return nil
	})

	if nil != err {
		log.Panicf("update the db of utxoset failed! %v\n", err)
	}
}

// 获取指定地址余额
func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	// 获取指定地址的UTXO
	UTXOS := utxoSet.FindUTXOWithAddress(address)
	fmt.Println("utxos : ", UTXOS)
	var amount int64 // 余额
	for _, utxo := range UTXOS {
		fmt.Printf("\tutxo-hash : %x\n", utxo.TxHash)
		fmt.Printf("\tutxo-Index : %d\n", utxo.Index)
		fmt.Printf("\tutxo-Ripemd160Hash : %x\n", utxo.Output.Ripemd160Hash)
		fmt.Printf("\tutxo-Value : %d\n", utxo.Output.Value)
		fmt.Println("----------------------------------------------------")

		amount += utxo.Output.Value
	}
	return amount
}
// 查找指定地址UTXO(utxo table)
func (utxoSet *UTXOSet) FindUTXOWithAddress(address string) []*UTXO {
	var utxos []*UTXO
	// 查找数据库的utxoTable表
	utxoSet.BlockChain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		if nil != b {
			// 游标
			c := b.Cursor()
			// 游标迭代
			// k -> 交易哈希
			// v -> 输出结构的字节数组
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTXOutputs(v)
				for _, utxo := range txOutputs.TxOutputs {
					if utxo.UnLockScriptPubkeyWithAddress(address) {
						utxo_signle := UTXO{Output:utxo}
						utxos = append(utxos, &utxo_signle)
					}
				}
			}
		}
		return nil
	})

	return utxos
}