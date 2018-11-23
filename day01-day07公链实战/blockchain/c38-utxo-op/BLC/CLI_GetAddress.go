package BLC

import "fmt"

func (cli *CLI) getAddresses() {
	fmt.Println("打印所有钱包地址...")
	wallets := NewWallets()
	for address, _ := range wallets.Wallets {
		fmt.Printf("address [%s]\n", address)
	}
}
