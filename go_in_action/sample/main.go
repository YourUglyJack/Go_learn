package main

import (
	_ "./matchers"
	"./search"
	"log"
	"os"
)

/*
程序的入口
*/

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定项做搜索
	search.Run("foxnews")
}
