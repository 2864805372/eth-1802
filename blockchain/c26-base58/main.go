package main

import (
	"blockchain/c26-base58/Blockchain"
	"fmt"
)

func main()  {
	b58Encode := Blockchain.Base58Encode([]byte("this is the example"))
	fmt.Printf("b58Encode : %s\n", b58Encode)
	b58Decode := Blockchain.Base58Decode([]byte("1nj2SLMErZakmBni8xhSXtimREn"))
	fmt.Printf("b58Decode : %s\n", b58Decode)
}
