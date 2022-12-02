package main

import "fmt"

// 结构体中，小写是私有属性，大写是public
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p // 返回person的地址
}

func main() {

	fmt.Println(person{name: "bob", age: 20})

	fmt.Println(person{"alice", 20})
	//fmt.Println(person{"fred"})
	fmt.Println(&person{"anna", 40})

	fmt.Println(newPerson("jon"))

	s := person{"sean", 50}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	sp.age = -1
	fmt.Println(sp.age, s.age) // 都会改变，因为sp是s的地址
}
