package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// var ip = flag.Int("flagname", 1234, "help message for flagname")
	/*
		var flagvar int
		func init() {
			flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
			// 将指定flag绑定到定义的变量上面
		}

		-flag		只支持bool
		-flag=x
		-flag x  // non-boolean flags only	只支持非bool类型
	*/

	flagPrintChain := flag.String("printchain","btc", "print the block chain")
	// name : flag名称
	// value : 默认值
	// usage : 用法

	flag.Usage()
	flag.Parse() // 解析
	fmt.Printf("the flag of block chain is : %s\n", *flagPrintChain)

	fmt.Printf("length : %d\n", len(os.Args))
}
