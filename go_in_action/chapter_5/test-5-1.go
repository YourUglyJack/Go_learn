package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func main() {
	var bill user // 当声明变量时会被初始化，要么为默认值0，要么为自定义的初始化值
	fmt.Println(bill)

	lisa := user{
		name:       "Lisa",
		email:      "lisa@gmail.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(lisa)

	jack := user{"Jack", "leety589589@gmail.com", 123, true}
	fmt.Println(jack)

	sysAdmin := admin{jack, "admin"}
	fmt.Println(sysAdmin)
}
