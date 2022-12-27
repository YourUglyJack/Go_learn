package main

import "../../zinx/znet"

func main() {
	// 创建一个server具柄，使用zinx的api
	s := znet.NewServer("[zinx v1]")
	// 启动server
	s.Server()
}
