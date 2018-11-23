package main

import (
	"blockchain/c02-blc/BLC"
	"fmt"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()

	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1,
		[]byte("a send 100 btc to b"),
		bc.Blocks[len(bc.Blocks) - 1].Hash)

	bc.AddBlock(bc.Blocks[len(bc.Blocks) - 1].Height + 1,
		[]byte("b send 10 btc to c"),
		bc.Blocks[len(bc.Blocks) - 1].Hash)

	for _, block := range bc.Blocks{
		fmt.Printf("block : %v\n", block)
	}
}