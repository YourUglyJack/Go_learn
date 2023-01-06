package ziface

// 路由接口，路由是使用者给该链接自定义的处理业务的方法
// 路由里面的Irequest包含了连接的信息和数据

type IRouter interface {
	PreHandle(req IRequest)
	Handle(req IRequest)
	PostHandle(req IRequest)
}
