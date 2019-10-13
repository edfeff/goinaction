package inserttype

import "fmt"

//嵌入类型-扩展或修改已有类型
func Show() {
	u := user{"wpp"}
	fmt.Println(u.Name())

	a := admin{user{"xxx"}, 11}
	fmt.Println(a.Name())
	fmt.Println(a.Age())

	invokeSayer(&u)
	invokeSayer(&a)

	invokeFucker(&u)
	invokeFucker(&a)

	//调用父类的方法
	a.user.fuck()
}

func invokeSayer(s Sayer) {
	fmt.Println("Sayer:", s.info())
}
func invokeFucker(f Fucker) {
	f.fuck()
}

//接口
type Sayer interface {
	info() string
}
type Fucker interface {
	fuck()
}

//基类
type user struct {
	name string
}

//user的方法
func (u *user) Name() string {
	return u.name
}

//user实现接口 admin也可以使用
func (u *user) info() string {
	return "user:" + u.name
}
func (u *user) fuck() {
	fmt.Println("user fuck", u.name)
}

//admin要重载fuck
func (a *admin) fuck() {
	fmt.Println("admin fuck", a.name)
}

//超类,继承user
type admin struct {
	user
	age int
}

func (a *admin) Age() int {
	return a.age
}
