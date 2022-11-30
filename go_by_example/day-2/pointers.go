package main

import "fmt"

func swap(a *int, b *int) {
	// a，b是地址，赋值的时候取一下值，不然就直接是赋地址了
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 1
	b := 2
	//The &i syntax gives the memory address of i, i.e. a pointer to i.
	swap(&a, &b)
	fmt.Println(a, b)
}
