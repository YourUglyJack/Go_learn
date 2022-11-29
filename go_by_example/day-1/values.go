package main

/*
Go has various value types including strings, integers, floats, booleans, etc
*/

import "fmt"

func main() {
	fmt.Println("Go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("5/2 =", 5/2)          // 5/2 = 2
	fmt.Println("7.0/3.0 = ", 7.0/3.0) // 7.0/3.0 =  2.3333333333333335

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
