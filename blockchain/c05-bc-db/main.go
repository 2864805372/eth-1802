package main

import (
	"blockchain/c05-bc-db/BLC"
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()

	bc.AddBlock([]byte("a send 100 btc to b"))

	bc.AddBlock([]byte("b send 10 btc to c"))

	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if nil != b {
			hash := b.Get([]byte("l"))
			fmt.Printf("value : %x\n", hash)
			block := b.Get(hash)
			fmt.Printf("heigth : %v\n", BLC.DeserializeBlock(block).Height)
		}else {
			fmt.Printf("the bucket is nil \n")
		}
		return nil
	})
}