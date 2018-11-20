package BLC

import (
	"bytes"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"net"
)

// 服务管理
// 3000 作为主节点地址
var knowNodes = []string{"localhost:3000"}

// 代表当前区块版本信息(决定是否需要同步)
type Version struct {
	Version		int		// 版本
	Height		int		// 当前节点的区块高度
	AddrFrom	string	// 当前节点地址
}
// 节点地址
var nodeAddress string
// 启动服务器
func startServer(nodeID string)  {
	fmt.Printf("启动服务[%s]...\n", nodeAddress)
	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	// 1. 监听节点
	listen, err := net.Listen(PROTOCOL,nodeAddress)
	if nil != err {
		log.Panicf("listen address of %s failed! %v\n", nodeAddress, err)
	}
	defer listen.Close()
	bc := BlockchainObject(nodeID)
	// 两个节点，主节点负责保存所有数据，钱包节点负责发送请求同步数据
	if nodeAddress != knowNodes[0] {
		// 非主节点，向主节点发送请求，同步数据
		// sendMessage(knowNodes[0], nodeAddress)
		sendVersion(knowNodes[0], bc)
	}
	// 2. 接收请求
	for {
		conn, err := listen.Accept()
		if nil != err {
			log.Panicf("accept connect failed! %v\n", err)
		}
		request, err := ioutil.ReadAll(conn)
		if nil != err {
			log.Panicf("Receive Message failed! %\n", err)
		}
		fmt.Printf("Receive a Message : %v\n", request)
		// 3. 处理请求

	}
}

// 节点发送请求
func sendMessage(to string, msg []byte)  {
	fmt.Println("向服务器发送请求")
	// 1. 连接服务器
	conn, err := net.Dial(PROTOCOL, to)
	if nil != err {
		log.Panicf("connect to server [%s] failed! %v\n", to, err)
	}
	defer conn.Close()
	// 要发送的数据添加到请求中
	_, err = io.Copy(conn, bytes.NewReader(msg))
	if nil != err {
		log.Panicf("add the data failed! %v\n", err)
	}
}
// 区块版本验证
func sendVersion(toAddress string, bc *BlockChain)  {
	// 在比特币中，消息是底层的比特序列
	// 前12个字节指定命令名称(version)
	// 后面的字节包含gob编码的相关消息结构
	heigth := 1
	// 得到数据
	data := gobEncode(Version{NODE_VERSION, heigth, nodeAddress})
	// 发送version命令
	request := append(commandToBytes(VERSION), data...)
	sendMessage(toAddress, request)
}