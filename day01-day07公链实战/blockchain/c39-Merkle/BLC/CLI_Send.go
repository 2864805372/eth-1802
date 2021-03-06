package BLC

import (
	"fmt"
	"os"
)

// 发送交易
func (cli *CLI) send(from, to, amount []string)  {
	// 检测数据库
	if !dbExists() {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	// 获取区块链对象
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	// 发起转账，产生挖矿
	blockchain.MineNewBlock(from, to, amount)

	// 更新utxo table
	utxoSet := &UTXOSet{blockchain}
	utxoSet.Update()
}