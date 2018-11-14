package BLC

// 区块链管理相关

// 区块链基本结构
type BlockChain struct {
	Blocks []*Block // 存储有序的区块
}

// 初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	// 生成创世区块
	genesisBlock  := CreateGenesisBlock("the init of the blockchain")
	return &BlockChain{[]*Block{genesisBlock}}
}

// 添加新的区块到区块链中
func (bc *BlockChain) AddBlock(height int64, data []byte, prevBlockHash []byte)  {
	newBlock := NewBlock(height,prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}