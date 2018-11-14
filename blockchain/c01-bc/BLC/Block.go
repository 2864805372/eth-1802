package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// 区块管理相关

// 实现一个最基本的区块结构
type Block struct {
	TimeStamp 			int64		// 区块时间戳，代表区块产生时间
	Height				int64		// 区块高度(索引、号码)，代表当前区块数量
	PrevBlockHash		[]byte		// 前区块哈希
	Hash 				[]byte		// 当前区块哈希
	Data				[]byte		// 交易数据
}

//创建新的区块
func NewBlock(height int64,prevBlockHash []byte,data []byte) *Block {
	var block Block
	block=Block{
		TimeStamp:time.Now().Unix(),
		Height:height,
		PrevBlockHash:prevBlockHash,
		Data:data,
	}
	// 通过计算生成当前区块的哈希
	block.SetHash()
	return &block
}

// 计算区块哈希
func (b *Block)SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	// 拼接区块的所有属性，进行哈希
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	},[]byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash=hash[:]
}