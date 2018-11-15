package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/labstack/gommon/log"
)

// 交易管理相关
type Transaction struct {
	// 交易的唯一标识符
	TxHash 	[]byte
	// 输入列表
	Vins	[]*TxInput
	// 输出列表
	Vouts	[]*TxOutput
}

// 生成交易哈希
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if nil != err {
		log.Panicf("tx hash encoded failed! %v\n", err)
	}
	// 生成交易哈希
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}
// 生成coinbase交易
/*
	address : 地址
*/
func NewCoinbaseTransaction(address string) *Transaction {
	// 输入
	txInput := &TxInput{[]byte{}, -1, "Genesis Block"}
	// 输出
	txOutput := &TxOutput{10, address}

	txCoinbase := Transaction{nil, []*TxInput{txInput}, []*TxOutput{txOutput}}

	// hash
	txCoinbase.HashTransaction()
	return &txCoinbase
}