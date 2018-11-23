package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	// sha256
	// 新建256哈希对象
	hasher := sha256.New()
	hasher.Write([]byte("eth1802"))
	bytes := hasher.Sum(nil)
	fmt.Printf("sha256:%x\n", bytes)

	// ripemd160
	hash160 := ripemd160.New()
	hash160.Write([]byte("eth1802"))
	bytesRipemd := hash160.Sum(nil)
	fmt.Printf("ripemd160:%x\n",bytesRipemd)
}
