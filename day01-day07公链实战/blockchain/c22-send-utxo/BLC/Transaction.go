package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
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

// 生成普通转账交易
func NewSimpleTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var txInputs []*TxInput 		// 输入
	var txOutputs []*TxOutput 		// 输出
	// 查找指定地址from的UTXO
	money, spendableUTXODic := bc.FindSpendableUTXO(from, int64(amount))
	fmt.Printf("money : %d\n", money)
	for txHash, indexArray := range spendableUTXODic {
		// input(消费源)
		txHashBytes, err :=  hex.DecodeString(txHash)
		if nil != err {
			log.Panicf("decode string %s failed! %v\n", err)
		}
		for _, index := range indexArray {
			txInput := &TxInput{txHashBytes, index, from}
			txInputs = append(txInputs, txInput)
		}
	}

	// 输出(转账源)
	txOutput := &TxOutput{int64(amount), to}
	txOutputs = append(txOutputs, txOutput)

	// 输出(找零)
	txOutput = &TxOutput{money - int64(amount), from}
	txOutputs = append(txOutputs, txOutput)
	tx := Transaction{nil, txInputs, txOutputs}
	// 生成新的交易哈希
	tx.HashTransaction()
	return &tx
}

// 判断指定交易是否是一个coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}