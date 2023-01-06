package main

import (
	"fmt"
	"net"
	"time"
)

const COUNT = 10

func main() {

	fmt.Println("Start client ...")

	// 1.链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	// 2.发送消息
	for i := 0; i < COUNT; i++ {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			fmt.Println("write err", err)
			return
		}
		// 3.接收服务器的回显
		buff := make([]byte, 512)
		cnt, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Read err", err)
			return
		}

		fmt.Printf("client recv info: %s, cnt:%d\n", buff, cnt)

		time.Sleep(1 * time.Second)
	}

}
