package znet

import (
	"../ziface"
	"fmt"
	"net"
)

type Connection struct {

	// 当前连接的socket Tcp套接字
	Conn *net.TCPConn
	// 连接的ID，也可以当成sessionID，全局唯一
	ConnID uint32
	// 是否关闭
	isClosed bool

	// todo 还是不理解这玩意
	Router ziface.IRouter

	ExitBuffChan chan bool

	//// 该连接处理方法的API 更新：从router里面拿
	//handleApi ziface.HandleFunc
	// 告知该连接关闭或退出的channel
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connId,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println(c.GetRemoteAddr().String(), "conn reader exit ...")
	defer c.Stop()

	for {
		buff := make([]byte, 512)
		_, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("recv buff err", err)
			c.ExitBuffChan <- true
			continue
		}
		// get data from client
		req := Request{
			conn: c,
			data: buff,
		}
		// 从路由中找到注册绑定Conn的对应的Handle
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
		//if err := c.handleApi(c.Conn, buff, cnt); err != nil {
		//	fmt.Println("connID", c.ConnId, "handle err", err)
		//	c.ExitBuffChan <- true
		//	continue
		//}
	}
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	go c.StartReader()
	for {
		select {
		case <-c.ExitBuffChan:
			// 得到消息就退出，不再阻塞
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// todo

	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Close conn err...", err)
		return
	}

	// 通知连接已关闭
	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}
