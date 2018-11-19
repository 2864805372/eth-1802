package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/ripemd160"
)

// 钱包相关
// 版本
const version = byte(0x00)

// checksum 长度
const addressCheckSumLen = 4

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

// 对公钥进行双哈希(sha256->ripemd160)
func Ripemd160Hash(pubKey []byte) []byte {
	// 1. sha256
	hash256 := sha256.New()
	hash256.Write(pubKey)
	hash := hash256.Sum(nil)

	// 2. ripemd160
	rmd160 := ripemd160.New()
	rmd160.Write(hash)
	return rmd160.Sum(nil)
}

// 通过钱包获取地址
func (w *Wallet) GetAddress() []byte {
	// 1. 获取160哈希结果
	ripemd160Hash := Ripemd160Hash(w.PublicKey)
	// 2. 生成version并加入到hash中
	version_ripemd160Hash := append([]byte{version}, ripemd160Hash...)
	// 3. 生成校验和数据
	checkSumBytes := CheckSum(version_ripemd160Hash)
	// 4 拼接校验和
	bytes := append(version_ripemd160Hash, checkSumBytes...)
	// 5. 调用base58Encode生成地址
	base58Bytes := Base58Encode(bytes)
	return base58Bytes
}

// 生成校验和
func CheckSum(payload []byte) []byte {
	first_hash := sha256.Sum256(payload)
	second_hash := sha256.Sum256(first_hash[:])
	return second_hash[:addressCheckSumLen] // 取4个字节
}

// 判断地址有效性
func IsValidForAddress(address []byte) bool {
	// 1. 地址通过base58Decode进行解码
	version_pubkey_checkSumBytes := Base58Decode(address) // 25位
	// 2. 拆分，进行校验和的检验
	checkSumBytes := version_pubkey_checkSumBytes[len(version_pubkey_checkSumBytes)-addressCheckSumLen:]
	versoin_ripemd160 := version_pubkey_checkSumBytes[:len(version_pubkey_checkSumBytes)-addressCheckSumLen]
	checkBytes := CheckSum(versoin_ripemd160)
	if bytes.Compare(checkSumBytes, checkBytes) == 0 {
		return true
	}
	return false
}