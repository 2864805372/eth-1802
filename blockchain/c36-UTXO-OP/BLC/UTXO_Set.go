package BLC

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
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

// 重置UTXO，可以在创建区块链的时候调用
func (utxoSet *UTXOSet)ResetUTXOSet()  {
	// 在第一创建的时候更新utxo table
	// 采用覆盖的方式，如果指定表原本存在，则先删除
	err := utxoSet.BlockChain.DB.Update(func(tx *bolt.Tx) error {
		// 查找UTXO表
		b := tx.Bucket([]byte(utxoTableName))
		if nil != b {
			tx.DeleteBucket([]byte(utxoTableName))
			c, _ := tx.CreateBucket([]byte(utxoTableName))
			if nil != c {
				// 查找所有未花费输出
				txOutputsMap := utxoSet.BlockChain.FindUTXOMap()
				for keyHash, output := range txOutputsMap {
					txHash, _ := hex.DecodeString(keyHash)
					// 存入utxo table
					err := c.Put(txHash, output.Serialize())
					if nil != err {
						log.Panicf("put the utxo into table failed! %v\n", err)
					}
				}
			}
		}
		return nil
	})

	if nil != err {
		log.Panicf("update the db of utxoset failed! %v\n", err)
	}
}