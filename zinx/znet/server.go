package znet

import (
	"../ziface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	// 服务器的名称
	Name string
	// 服务器所绑定的ip版本
	IPVersion string
	// 服务器所监听的ip
	IP string
	// 服务器所监听的端口
	Port int
}

// 实现ziface.IServer的全部接口方法

// Start 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[Start] Server which listens at %s:%d is starting...\n", s.IP, s.Port)

	// 开启一个goroutine去监听
	go func() {
		// 1 获取tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve tcp addr err", err)
			return
		}
		// 2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s.IPVersion, "err", err)
		}

		fmt.Println("start zinx server", s.Name, " succ, now start to listen...")

		// 3启动server网络链接业务
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// todo 设置服务器最大链接控制，超过最大链接就关闭
			// todo 处理新链接请求的业务方法，此时handler和conn是绑定的

			// 初版暂定echo业务
			go func() {
				// 等待客户端发送数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Receive buff err", err)
					}
					// echo
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write buff err", err)
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[Stop] server, ", s.Name)

	// todo 将链接信息或资源清楚或关闭
}

func (s *Server) Server() {
	s.Start()

	// todo 在启动服务后，可以做些额外的业务

	for {
		time.Sleep(10 * time.Second)
	}
}

// NewServer 创建服务器句柄
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
