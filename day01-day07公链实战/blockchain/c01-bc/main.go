package main

import (
	"blockchain/c01-bc/BLC"
	"fmt"
)

func main() {
	block := BLC.NewBlock(1,nil,[]byte("this is the genesis block"))
	fmt.Printf("block : %x\n", block.Hash)
}