package BLC

import "fmt"

// 查询余额
func (cli *CLI) getBalance (from string)  {
	// 查找指定地址UTXO

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	utxoSet := &UTXOSet{blockchain}
	amount := utxoSet.GetBalance(from)
	fmt.Printf("\t地址 [%s] 的余额为 [%d]\n", from, amount)
}