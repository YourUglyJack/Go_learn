package main

import (
	"../../zinx/ziface"
	"../../zinx/znet"
	"fmt"
)

// PingRouter 自定义路由
type PingRouter struct {
	znet.BaseRouter // 继承基础的路由基类
}

// PreHandle 测试PreHandle
func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	s := req.GetConnection().GetTCPConnection()
	_, err := s.Write([]byte("before ping.....\n"))
	if err != nil {
		fmt.Println("Writer err...", err)
		return
	}
}

func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle")
	s := req.GetConnection().GetTCPConnection()
	_, err := s.Write([]byte("ping...ping..ping...\n"))
	if err != nil {
		fmt.Println("Writer err...", err)
		return
	}
}

func (pr *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	s := req.GetConnection().GetTCPConnection()
	_, err := s.Write([]byte("after ping....\n"))
	if err != nil {
		fmt.Println("Writer err...", err)
		return
	}
}

func main() {
	s := znet.NewServer("[zinx v0.3]")
	s.AddRouter(&PingRouter{})
	s.Server()
}
