package BLC

// 创建区块链
func (cli *CLI) createBlockchainWithGenesis(address string) {
	CreateBlockChainWithGenesisBlock(address)
}