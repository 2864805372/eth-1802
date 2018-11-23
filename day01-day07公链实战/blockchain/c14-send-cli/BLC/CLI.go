package BLC

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

// CLI结构
type CLI struct {
	BC *BlockChain 
}

// 用法展示
func PrintUsage()  {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblockchain -address address -- 创建区块链\n")
	fmt.Printf("\taddblock -data DATA -- 交易数据\n")
	fmt.Printf("\tprintchain -- 输出区块链信息\n")
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT -- 发起转账交易\n")
}

// 检测参数数量
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		// 如果参数数量不对，直接退出程序
		os.Exit(1)
	}
}

// 发送交易
func (cli *CLI) send()  {

}
// 添加区块
func (cli *CLI) addBlock(txs []*Transaction) {
	if !dbExists() {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockchain := BlockchainObject() // 获取区块链对象
	blockchain.AddBlock(txs)
}
// 输出区块链信息
func (cli *CLI) printchain() {
	if !dbExists() {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockchain := BlockchainObject() // 获取区块链对象
	blockchain.PrintChain()
}
// 创建区块链
func (cli *CLI) createBlockchainWithGenesis(address string) {
	CreateBlockChainWithGenesisBlock(address)
}

// 命令行运行函数
func (cli *CLI)Run()  {
	// 1. 检测参数数量
	IsValidArgs()
	// 2. 新建命令
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBLCWithGenesisCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	// 发送交易
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	// 3. 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC to everyone","交易数据")
	flagCreateBlockchainArg := createBLCWithGenesisCmd.String("address","","the address of create blockchain")
	// 转账命令行参数
	flagFromArg := sendCmd.String("from", "", "转账源地址...")
	flagToArg := sendCmd.String("to", "", "转账目标地址...")
	flagAmount := sendCmd.String("amount", "", "转账金额...")
	switch os.Args[1] {
	case "send":
		if err := sendCmd.Parse(os.Args[2:]); nil != err {
			if nil != err {
				log.Panicf("parse cmd of send failed! %v\n", err)
			}
		}
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse cmd of add block failed! %v\n", err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse cmd of printchain failed! %v\n", err)
		}
	case "createblockchain":
		if err := createBLCWithGenesisCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse cmd of create block chain failed! %v\n", err)
		}
	default:
		PrintUsage()
		os.Exit(1)
	}
	// 添加转账命令
	if sendCmd.Parsed() {
		if *flagFromArg == "" {
			fmt.Println("源地址不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		if *flagToArg == "" {
			fmt.Println("目标地址不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		if *flagAmount == "" {
			fmt.Println("金额不能为空...")
			PrintUsage()
			os.Exit(1)
		}
		fmt.Printf("\tFROM:[%s]\n", *flagFromArg)
		fmt.Printf("\tTO:[%s]\n", *flagToArg)
		fmt.Printf("\tAMOUNT:[%s]\n", *flagAmount)
	}
	// 添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}

	// 输出区块链信息命令
	if printChainCmd.Parsed() {
		cli.printchain()
	}

	// 创建区块链
	if createBLCWithGenesisCmd.Parsed() {
		cli.createBlockchainWithGenesis(*flagCreateBlockchainArg)
	}
}