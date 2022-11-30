package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 1, 1
}

// variadic functions 可变参数类型  eg: fmt.Println
func sum(nums ...int) {
	// within function, the type of nums is equivalent to []int
	fmt.Print(nums, " total: ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 1) // 新变量用:= 若之后再更改就直接 =
	fmt.Println("1 + 1 =", res)

	res = plus(2, 2)
	fmt.Println("2 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3, 4)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
