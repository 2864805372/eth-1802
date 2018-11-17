package main

import (
	"blockchain/c29-wallets/Blockchain"
	"fmt"
)

func main()  {
	// 创建wallet
	//wallet := Blockchain.NewWallet()
	//address := wallet.GetAddress()
	//fmt.Printf("address : %s\n", address)
	//fmt.Printf("validation of address %s is %v\n", address,
	//	Blockchain.IsValidForAddress(address))
	wallets := Blockchain.NewWallets()
	wallets.CreateWallet()
	wallets.CreateWallet()
	fmt.Printf("wallets : %v\n", wallets.Wallets)
}
