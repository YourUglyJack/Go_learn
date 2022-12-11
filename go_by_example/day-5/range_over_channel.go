package main

import "fmt"

func main() {
	quene := make(chan string, 2)
	quene <- "un"
	quene <- "deux"
	close(quene) // 记得把通道关了，不然会报错，因为通道一直等待value
	for elem := range quene {
		fmt.Println(elem)
	}
}
