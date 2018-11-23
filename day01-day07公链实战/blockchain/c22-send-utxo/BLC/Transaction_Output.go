package BLC

// 交易输出
type TxOutput struct {
	// 金额
	Value 			int64
	// 用户名(该UTXO的拥有者)
	ScriptPubkey	string
}

// output身份验证
func (txOutput *TxOutput) UnLockScriptPubkeyWithAddress(address string) bool {
	return address == txOutput.ScriptPubkey
}