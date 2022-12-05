package main

import "fmt"

// Duration 是一种描述时间间隔的类型，单位是纳秒（ns）。这个类型使用内置的int64 类型作为其表示
type Duration int64

func main() {
	var dur Duration
	// dur = int64(1000)  是错误的
	//在Duration 类型的声明中，我们把int64 类型叫作Duration 的基础类型。不过，虽然int64 是基础 类型，Go 并不认为Duration 和int64 是同一种类型。
	fmt.Println(dur)
}
