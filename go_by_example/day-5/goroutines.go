package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("start a anonymous goroutine")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
