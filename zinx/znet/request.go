package znet

import "../ziface"

type Request struct {
	conn ziface.IConnection // zinx封装的连接
	data []byte
}

// GetConnection 获取请求的连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 相当于get方法
func (r *Request) GetData() []byte {
	return r.data
}
