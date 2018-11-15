package BLC

import (
	"bytes"
	"encoding/binary"
	"github.com/labstack/gommon/log"
)

// int64转换成字节数组
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer,binary.BigEndian,data)
	if nil != err {
		log.Panicf("int to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}