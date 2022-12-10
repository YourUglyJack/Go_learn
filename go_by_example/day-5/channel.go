package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	// The <-channel syntax receives a value from the channel
	msg := <-message // msg并不是通道，而是通道传来的值

	// message type: chan string
	// msg type: string
	// msg: ping
	fmt.Println("message type:", reflect.TypeOf(message))
	fmt.Println("msg type:", reflect.TypeOf(msg))
	fmt.Println("msg:", msg)
}
