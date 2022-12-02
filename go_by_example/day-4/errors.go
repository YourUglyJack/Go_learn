package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	pro string
}

// It’s possible to use custom types as errors by implementing the Error() method on them.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.pro)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, pro: " can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// f1 内置error
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	// f2 自定义error
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42) // e是自定义error的地址
	// todo wtf: ae, ok := e.(*argError) ???
	// 断言机制，判断e是否是argError类型，ok就是 是不是这种类型
	if ae, ok := e.(*argError); ok {
		fmt.Print(ae.arg)
		fmt.Println(ae.pro)
	}
}
