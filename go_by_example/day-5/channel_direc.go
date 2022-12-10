package main

import (
	"fmt"
	"reflect"
)

func ping(pings chan<- string, msg string) {
	//  a channel for sending values
	pings <- msg
	fmt.Println("type of pings", reflect.TypeOf(pings))
}

func pong(pings <-chan string, pongs chan<- string) {
	// pings 是接收value的chan，pong是发送value的chan???
	fmt.Println("type of pings", reflect.TypeOf(pings))
	//fmt.Println("type of pongs", reflect.TypeOf(pongs))
	msg := <-pings // pings是某通道传来的值
	fmt.Println("type of msg", reflect.TypeOf(msg))
	pongs <- msg // 将msg传给pongs通道
	fmt.Println("type of pongs", reflect.TypeOf(pongs))
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
