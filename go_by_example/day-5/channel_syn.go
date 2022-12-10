package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second * 20)
	fmt.Println("Done...")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
