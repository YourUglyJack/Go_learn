package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	// 格式化字符串并返回该字符串
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "hello word",
	}

	fmt.Printf("co={num:%v, str:%s}\n", co.base.num, co.str)

	fmt.Println(co.describe())
	fmt.Println(co.base.describe())

	type describe interface {
		describe() string
	}

	var d describe = co // 声明一个接口
	fmt.Println(d)
}
