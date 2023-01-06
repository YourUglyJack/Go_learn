package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client test start.....")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Create conn err...", err)
		return
	}

	for i := 0; i <= 3; i++ {
		_, err := conn.Write([]byte("Zinx v0.3..."))
		if err != nil {
			fmt.Println("Writer err...", err)
			return
		}

		buff := make([]byte, 8192)
		cnt, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Read err...", err)
			return
		}

		fmt.Printf("Server callback: %s, cnt:%d\n", buff, cnt)
		fmt.Println("--------------------------------------------")
		time.Sleep(1 * time.Second)
	}
}
