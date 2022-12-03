package search

import (
	"encoding/json"
	"os"
)

/*
用于读取json文件
*/

// 常量名使用小写，说明只能包内访问，包外不能访问
const dataFile = "D:\\code\\GoLang\\go_learn\\Go_learn\\go_in_action\\sample\\data\\data.json"

// Feed 包含我们需要处理的数据源信息
type Feed struct { // 名为Feed的类型结构
	// 每个标记将结构类型里的字段对应到JSON文档里指定名字的字段
	Name string `json:"site"` // `` 标记(tag) 描述了JSON解码的元数据
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件，返回切片与异常，切片里的每个元素都是指向Feed类型的指针，error可以用来判断函数是否调用成功
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil { // 打开文件异常
		return nil, err
	}

	// 当函数返回时才会执行，保障就算函数崩溃了也会关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 改切片的每一项是一个指向Feed类型的指针
	var feeds []*Feed // 声明了一个名为feeds，值为nil的切片
	// 我们使用之前调用Open 返回的文件句柄调用NewDecoder 函数，并得到一个
	// 指向Decoder 类型的值的指针。之后再调用这个指针的Decode 方法，传入切片的地址。之后
	// Decode 方法会解码数据文件，并将解码后的值以Feed 类型值的形式存入切片里。
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}
