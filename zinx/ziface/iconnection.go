package ziface

import "net"

// IConnection 定义连接接口，抽象层
type IConnection interface {

	// Start 启动连接
	Start()

	// Stop 关闭连接
	Stop()

	// GetTCPConnection 从连接获取socket
	GetTCPConnection() *net.TCPConn

	// GetConnId 获取连接的ID
	GetConnId() uint32

	// GetRemoveAddr 获取客户端地址
	GetRemoveAddr() net.TCPAddr
}

// HandleFunc 定义一个统一处理连接业务的接口，第一个参数是原生socket 第二个参数是客户端请求的数据 第三个参数是客户端请求数据的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
