package BLC

type TxInput struct {
	// 交易哈希(不是当前交易的哈希，而是引入的上一笔交易的哈希)
	TxHash 		[]byte
	// 引用的上一笔交易的输出的索引
	Vout		int
	// 用户名
	ScriptSig	string
}