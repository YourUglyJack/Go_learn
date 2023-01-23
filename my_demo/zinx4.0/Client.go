package main

import (
	"../../zinx/znet"
	"fmt"
	"net"
)

func main() {
	// 客户端goroutine，负责模拟粘包数据，然后进行发送
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Dial err...", err)
		return
	}

	// 创建一个封包拆包对象
	dp := znet.NewDataPack()

	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("Pack msg1 err...", err)
		return
	}

	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("Pack msg2 err...", err)
		return
	}

	// 将sendData1和2写到一起，组成一个粘包
	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	// 客户端阻塞
	select {}
}
