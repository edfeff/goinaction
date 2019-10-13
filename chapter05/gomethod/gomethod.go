package gomethod

import "fmt"

//go 方法
func Show() {
	p := Person{"wpp", "wpp@e.com"}
	notify(p)
	p.pnotify()
	p.changeEmail("wpp@11.com")
	p.pnotify()
}

type Person struct {
	name  string
	email string
}

//普通函数
func notify(p Person) {
	fmt.Printf("name:%s,email:%s\n", p.name, p.email)
}

//针对结构体的方法
//(p Person)称为 值接收者
//当结构体调用此方法时，此方法会创建出一个结构体的副本进行操作
func (p Person) pnotify() {
	fmt.Printf("name:%s,email:%s\n", p.name, p.email)
}

//传递结构体的指针，可以进行值的修改
// (p *Person)称为指针接受者
//当结构体调用此方法时，直接操作了原有的结构体
func (p *Person) changeEmail(email string) {
	p.email = email
}
