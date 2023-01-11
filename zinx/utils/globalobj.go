package utils

import (
	"../ziface"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
存储一切有关zinx框架的全局参数，供其他模块使用
一些参数也可以通过，用户根据zinx.json
*/

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	Version   string

	MaxPacketSize uint32 // 数据包到最大值
	MaxConn       int
}

// define a global obj
var GlobalObject *GlobalObj

// Reload 读取用户配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("Reload err", err)
		return
	}
	// 将json数据解析到全局配置对象中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}

}

func init() {
	// 初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:          "ZinxServer",
		Version:       "v0.4",
		TcpPort:       7777,
		Host:          "127.0.0.1",
		MaxConn:       10,
		MaxPacketSize: 8192,
	}
	// 从配置文件，读取用户配置的参数，可覆盖默认值
	GlobalObject.Reload()
}
