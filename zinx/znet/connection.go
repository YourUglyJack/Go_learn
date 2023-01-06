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
	ConnId uint32
	// 是否关闭
	isClosed bool

	// 该连接处理方法的API
	handleApi ziface.HandleFunc
	// 告知该连接关闭或退出的channel
	ExitBuffChan chan bool
}

func NewConnetion(conn *net.TCPConn, connId uint32, callback ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnId:       connId,
		isClosed:     false,
		handleApi:    callback,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println(c.GetRemoteAddr().String(), "conn reader eixt ...")
	defer c.Stop()

	for {
		buff := make([]byte, 512)
		cnt, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("recv buff err", err)
			c.ExitBuffChan <- true
			continue
		}
		if err := c.handleApi(c.Conn, buff, cnt); err != nil {
			fmt.Println("connID", c.ConnId, "handle err", err)
			c.ExitBuffChan <- true
			continue
		}
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

	// 通知连接已关闭
	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}
