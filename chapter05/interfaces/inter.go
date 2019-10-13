package interfaces

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name string
}
type admin struct {
	name string
}

//user 实现了notify接口
func (u *user) notify() {
	fmt.Println("user", u.name)
}
func (a *admin) notify() {
	fmt.Println("admin:", a.name)
}

//方法参数为接口
func sendNotification(n notifier) {
	fmt.Println("接口调用")
	n.notify()
}
func Show() {
	u1 := user{"wpp"}
	a1 := admin{"wpp"}

	//fmt.Println(u1)
	//不能传入u1，因为user没有实现notifier接口
	//sendNotification(u1)
	//只有传入地址
	//多态
	sendNotification(&u1)
	sendNotification(&a1)
}

//
