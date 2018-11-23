package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

// 标准JSON格式转切片
// 在windows下，JSON转账成slice的标准的输入格式：
// bc.exe send -from "[\"Alice\"]" -to "[\"Bob\"]" -amount "[\"2\"]"
func JSONToSlice(jsonString string) []string {
	var strSlice []string
	// 通过json包进行转换
	if err := json.Unmarshal([]byte(jsonString), &strSlice); err != nil {
		log.Panicf("json to []string failed! %v\n", err)
	}
	return strSlice
}