package gotype

import "fmt"

//go的类型系统
func Show() {
	//userDefineType()
	changeTypeName()
}

//自定义结构体
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

//结构体中使用结构体
type admin struct {
	person user
	level  string
}

// 用户自定义类型
func userDefineType() {
	//使用零值初始化
	var u2 user
	fmt.Println(u2)

	//赋值初始化
	wpp := user{
		name:       "wpp",
		email:      "edfeff@163.com",
		ext:        0,
		privileged: true,
	}
	fmt.Println(wpp)

	// 一行初始化
	u3 := user{"u3", "u3@qq.com", 19, false}
	fmt.Println(u3)

	//定义一个admin
	ad1 := admin{
		person: user{"ad1", "ad@11.com", 100, true},
		level:  "level 1",
	}
	fmt.Println(ad1)
}

//类型改名
type Integer int
type Admin admin
type User user

func changeTypeName() {
	var i Integer
	i = 10
	fmt.Println(i)
	var user1 User
	var admin1 Admin
	fmt.Println(user1)
	fmt.Println(admin1)
}
