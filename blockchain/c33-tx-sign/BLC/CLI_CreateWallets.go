package BLC

import "fmt"

// 通过命令行实现钱包创建
func (cli *CLI) CreateWallets() {
	// 创建一个集合对象
	// 在钱包文件已存在的情况下，把文件中原有的数据全部读取出来
	// 保存到wallets对象中
	wallets := NewWallets()
	wallets.CreateWallet()
	fmt.Printf("wallets : %v\n", wallets)
}