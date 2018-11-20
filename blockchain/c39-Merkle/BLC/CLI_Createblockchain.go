package BLC

// 创建区块链
func (cli *CLI) createBlockchainWithGenesis(address string) {
	blockchain := CreateBlockChainWithGenesisBlock(address)
	defer blockchain.DB.Close()
	// 设置utxoSet操作
	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}