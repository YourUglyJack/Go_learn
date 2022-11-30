package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	fmt.Println(fact(3))

	// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	// 闭包也可以递归，但是在定义前，用var声明
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(3))
}
