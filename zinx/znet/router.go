package znet

import "../ziface"

// 路由基类，继承这个基类，如果有需要，再重写相应的方法
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req ziface.IRequest) {

}

func (br *BaseRouter) Handle(req ziface.IRequest) {

}

func (br *BaseRouter) PostHandle(req ziface.IRequest) {

}
