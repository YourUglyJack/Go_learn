package main

import "fmt"

type person struct {
	name  string
	email string
}

/*
关键字func 和函数名之间的参数被称作接收者，将函数与接收者的类型绑在一起。
如果一个函数有接收者，这个函数就被称为方法。
Go语言有两种接收者：值接收者和指针接收者
*/

// notify 使用：值接受者实现了一个方法
// 如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行。
func (p person) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n\n", p.name, p.email)
}

// changeEmail 使用：指针接受者实现了一个方法
// 当调用使用指针接收者声明的方法时，这个方法会共享调用方法时接收者所指向的值，
func (p *person) changeEmail(email string) {
	p.email = email
}

func main() {
	bill := person{"bill", "bill@gmail.com"}
	bill.notify() // notify接收的是bill的副本

	lisa := &person{name: "lisa", email: "lisa@gmail.com"}
	lisa.notify() // (*lisa).notify() go会自动调整
	lisa.changeEmail("lisa@email.com")
	lisa.notify()
}
