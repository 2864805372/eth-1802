package BLC

// 交易管理相关

type Transaction struct {
	// 交易的唯一标识符
	TxHash 	[]byte

	// 输入
	Vins	[]*TxInput
	// 输出
	Vounts	[]*TxOutput

}