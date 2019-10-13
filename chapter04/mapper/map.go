package mapper

import "fmt"

//go 的数据结构 map
//Show 映射类型
//
func Show() {
	//创建映射
	fmt.Println("创建映射")
	//创建 键是string 值是 int
	d1 := make(map[string]int)
	fmt.Println(d1)
	//创建 键和值都是string
	d2 := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	d2["black"] = "#FFFFFF"
	d2["write"] = "#000000"

	fmt.Println(d2)
	//映射的键不能是 切片、函数、和包含切片的结构体
	//d3 := map[[]string]string{}
	fmt.Println("使用映射")
	//nil映射 不能直接赋值
	//var d3 map[string]string
	//d3["a"] = "A"

	//判断存在
	//两个值时，第一个是值，第二个是存在标志
	value, exists := d2["red"]
	if exists {
		fmt.Println(value)
	}
	value = d2["green"]
	fmt.Println(value)
	//不存在时，go会返回该类型对应的零值，对于string，零值就是 “”
	value = d2["XXX"]
	fmt.Println(value)
	//迭代映射
	for k, v := range d2 {
		fmt.Printf("key:%s value:%s \n", k, v)
	}
	//删除映射中的一项
	delete(d2, "red")
	fmt.Println()
	for k, v := range d2 {
		fmt.Printf("key:%s value:%s \n", k, v)
	}
	//函数间传递映射，对映射的操作会互相影响
	f1()
}

func f1() {
	d2 := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	f2(d2)
	for key, _ := range d2 {
		fmt.Println(key)
	}
}

func f2(m map[string]string) map[string]string {
	_, b := m["red"]
	if b {
		delete(m, "red")
	}
	return m
}
