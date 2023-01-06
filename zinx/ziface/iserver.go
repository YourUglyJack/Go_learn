package ziface

// IServer 抽象接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Server 运行服务器
	Server()
	// AddRouter 给当前服务器注册一个路由业务的方法，供客户端连接处理使用
	AddRouter(router IRouter)
}
