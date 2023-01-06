package ziface

// 抽象层

type IRequest interface {
	GetConnection() IConnection // 获取请求的连接
	GetData() []byte            // 获取请求的数据
}
