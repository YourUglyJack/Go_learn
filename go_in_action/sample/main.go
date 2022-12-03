package sample

import (
	_ "./search"
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
