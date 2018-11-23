package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	msg := "this is the eg of base64 encode"
	// 编码
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	// dGhpcyBpcyB0aGUgZWcgb2YgYmFzZTY0IGVuY29kZQ==
	// dGhpcyBpcyB0aGUgZWcgb2YgYmFzZTY0IGVuY29kZQ==

	// 解码
	decoded, err := base64.StdEncoding.DecodeString("dGhpcyBpcyB0aGUgZWcgb" +
		"2YgYmFzZTY0IGVuY29kZQ==")
	if nil != err {
		panic(err)
	}
	fmt.Println(string(decoded))
}
