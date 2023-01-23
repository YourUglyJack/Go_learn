package main

import (
	"../../zinx/znet"
	"fmt"
	"io"
	"net"
)

// 测试拆包封包

func main() {
	// 创建tcp server
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Listen err...", err)
		return
	}
	fmt.Println("Start listen...")
	// 创建服务器goroutine，负责从客户端goroutine读取粘包数据并解析
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err...")
		}
		fmt.Println("Accept a Conn")
		// 处理客户端请求，可能有多个客户端同时发来，所以开goroutine去处理
		go func(conn net.Conn) {
			// 创建封包拆包的对象
			dp := znet.NewDataPack() // 为每个请求创建一个拆包封包对象
			for {
				// 读取流中的head部分
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData) // 会把headData填满
				if err != nil {
					fmt.Println("ReadFull err...")
					break
				}
				// 将headData字节流拆包到msg里
				msgHead, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("Unpack err...", err)
					return
				}
				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen()) // 在头部已经传输了这次Data到长度

					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("msg.Data ReadFull err...", err)
						return
					}

					fmt.Println("===> Recv Msg: ID=", msg.Id, "len=", msg.DataLen, "Data:", string(msg.Data))
				}
			}
		}(conn)
	}
}
