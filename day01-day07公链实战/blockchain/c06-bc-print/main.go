package main

import (
	"blockchain/c06-bc-print/BLC"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()

	//bc.AddBlock([]byte("a send 100 btc to b"))
	//
	//bc.AddBlock([]byte("b send 10 btc to c"))

	bc.PrintChain()
}