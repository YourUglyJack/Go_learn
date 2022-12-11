package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // block 2s
	fmt.Println("Timer 1 fired ")

	timer2 := time.NewTimer(time.Second)
	go func() { // start a goroutine to wait block 1s
		<-timer2.C
		fmt.Println(" Timer 2 fired ")
	}()
	stop2 := timer2.Stop() // main 线程在timer2 block等待结束之前就结束了timer2
	if stop2 {
		fmt.Println("Timer 2 stoped")
	}
	time.Sleep(2 * time.Second)
}
