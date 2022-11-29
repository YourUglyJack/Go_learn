package main

// Go supports constants of character, string, boolean, and numeric values.

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d) // 6e+11

	fmt.Println(int64(d)) // 600000000000

	fmt.Println(math.Sin(n))
}
