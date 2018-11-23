package main

import (
	"blockchain/c28-wallet-address/Blockchain"
	"fmt"
)

func main()  {
	//b58Encode := Blockchain.Base58Encode([]byte("this is the example"))
	//fmt.Printf("b58Encode : %s\n", b58Encode)
	//b58Decode := Blockchain.Base58Decode([]byte("1nj2SLMErZakmBni8xhSXtimREn"))
	//fmt.Printf("b58Decode : %s\n", b58Decode)
	// 创建wallet
	wallet := Blockchain.NewWallet()
	address := wallet.GetAddress()
	fmt.Printf("address : %s\n", address)
	fmt.Printf("validation of address %s is %v\n", address,
		Blockchain.IsValidForAddress(address))
}
