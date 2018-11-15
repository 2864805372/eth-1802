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
	fmt.Printf("\tcreateblockchain -- 创建区块链\n")
	fmt.Printf("\taddblock -data DATA -- 发起交易\n")
	fmt.Printf("\tprintchain -- 输出区块链信息\n")
}

// 检测参数数量
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		// 如果参数数量不对，直接退出程序
		os.Exit(1)
	}
}

// 添加区块
func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock([]byte(data))
}
// 输出区块链信息
func (cli *CLI) printchain() {
	cli.BC.PrintChain()
}
// 创建区块链
func (cli *CLI) createBlockchainWithGenesis() {
	CreateBlockChainWithGenesisBlock()
}

// 命令行运行函数
func (cli *CLI)Run()  {
	// 1. 检测参数数量
	IsValidArgs()
	// 2. 新建命令
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBLCWithGenesisCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	// 3. 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC to everyone","交易数据")

	switch os.Args[1] {
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

	// 添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)
	}

	// 输出区块链信息命令
	if printChainCmd.Parsed() {
		cli.printchain()
	}

	// 创建区块链
	if createBLCWithGenesisCmd.Parsed() {
		cli.createBlockchainWithGenesis()
	}
}