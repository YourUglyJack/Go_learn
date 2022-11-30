package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	fmt.Println("len(m):", len(m))

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v2 := m["k2"]
	fmt.Println("v2:", v2)

	delete(m, "k1")

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	// go判断key存在的方式！！！！
	_, prs := m["k2"]
	fmt.Println("m has k2?", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
