package Blockchain

// 钱包集合，对钱包进行维护、管理的文件

// 钱包集合的结构
type Wallets struct {
	// key : 地址
	// value : 钱包结构
	Wallets map[string] *Wallet
}

// 初始化钱包集合
func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string] *Wallet)
	return wallets
}

// 创建新的钱包,并且添加到钱包集合中
func (wallets *Wallets) CreateWallet() {
	// 新建钱包对象
	wallet := NewWallet()
	wallets.Wallets[string(wallet.GetAddress())] = wallet
}