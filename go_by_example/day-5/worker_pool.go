package main

import (
	"fmt"
	"time"
)

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // iter jobs通道里的值，如果没有就阻塞，除非通道关闭
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 // 把结果传入results通道
	}

}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ { // 开启三个goroutine，分别对应三个worker
		go workerPool(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j // 把工作编号写入jobs通道
	}

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.

	//for a := 1; a <= numJobs; a++ {
	//	fmt.Println("res:", <-results)
	//}
	close(jobs)

}
