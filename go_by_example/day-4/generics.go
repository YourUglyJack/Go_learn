package main

import "fmt"

// 泛型用大写？

// returns a slice of its keys
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// 链表的一个结点：指向下一个结点的指针+当前结点的值
type element[T any] struct {
	next  *element[T] // 指向下一个结点的指针
	value T           // 当前结点的值
}

// todo 不是很能看懂 List[T any]
type List[T any] struct {
	head, tail *element[T] // 头尾指针
}

// 尾插
func (lst *List[T]) Add(v T) {
	if lst.tail == nil { // 如果是空链表
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next // 更新尾指针
	}
}

// 遍历数组
func (lst *List[T]) getAll() []T {
	var elems []T // T类型的数组，名为elems，这个T类似泛型？
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {

	var m = map[int]string{1: "un", 2: "deux", 3: "trois"}
	fmt.Println("keys:", MapKeys(m))

	res := MapKeys[int, string](m)
	fmt.Println("keys:", res)

	lst := List[int]{}
	lst.Add(10)
	lst.Add(100)
	lst.Add(1000)
	fmt.Println("list:", lst.getAll())
}
