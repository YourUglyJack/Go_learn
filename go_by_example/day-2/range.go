package main

import "fmt"

/*
range iterates over elements in a variety of data structures
*/

func main() {

	nums := []int{2, 3, 4}
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("idx:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "abc" {
		// 输出的是asc ii:97,98,99
		fmt.Println(i, c)
	}

}
