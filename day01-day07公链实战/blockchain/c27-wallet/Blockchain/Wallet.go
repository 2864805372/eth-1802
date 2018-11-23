package Blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/labstack/gommon/log"
)

// 钱包相关

// 钱包结构
type Wallet struct {
	// 1. 私钥
	PrivateKey ecdsa.PrivateKey
	// 2. 公钥
	PublicKey 	[]byte
}

// 创建一个钱包
func NewWallet() *Wallet {
	// 获取公钥-私钥对
	privateKey, publicKey := newKeyPair()
	return &Wallet{PrivateKey:privateKey, PublicKey:publicKey}
}

// 生成公钥-私钥对
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// 获取椭圆
	curve := elliptic.P256()
	// 椭圆加密
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if nil != err {
		log.Panicf("ecdsa generate key failed! %v\n")
	}
	// 生成公钥
	pubKey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	return *priv, pubKey
}