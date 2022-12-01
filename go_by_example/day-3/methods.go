package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return (r.width + r.height) * 2
}

func printInfo(r rect) {
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
}
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.height, r.height)
	printInfo(r)

	rp := &r
	fmt.Println(r.height, r.height)
	printInfo(*rp)
}
