package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty s", s)

	fmt.Println("Now,start to set s")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after set:", s)

	fmt.Println("len(s) = ", len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s) // func copy(dst, src []Type) int  deepcopy!!!
	fmt.Println(c)
	fmt.Println("test:")
	c[0] = "q"
	fmt.Println("s:", s, " c", c)

	l := s[2:5]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2", l)

	l = s[2:]
	fmt.Println("slice 3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
