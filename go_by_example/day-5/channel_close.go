package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for { // 死循环？
			j, more := <-jobs // j是通道传来的值
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return // 退出死循环
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	// Closing a channel indicates that no more values will be sent on it.
	// This can be useful to communicate completion to the channel’s receivers.
	close(jobs) // 通道被关闭了，more就是false了
	fmt.Println("send all jobs")
	<-done
}
