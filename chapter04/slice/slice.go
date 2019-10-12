package slice

import "fmt"

//go 切片
// 3个属性
// 底层数组的指针  当前切片的元素个数   容量
func Show() {
	fmt.Println("切片")
	// 创建切片
	fmt.Println("创建切片")

	//类型 长度等于容量
	s1 := make([]string, 5)
	fmt.Println(s1)

	//长度为3 容量为5
	s2 := make([]int, 3, 5)
	fmt.Println(s2)

	//创建并初始化
	s3 := []string{"a", "b", "c"}
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s3)
	fmt.Println(s4)

	//使用索引声明
	s5 := []string{8: "a"}
	fmt.Println(s5)

	//数组和切片的声明区别
	//数组必须指定长度
	array := [3]int{1, 2, 3}
	//切片不用指定长度
	s6 := []int{1, 2, 3}
	fmt.Println(array)
	fmt.Println(s6)

	//nil 和 空切片

	//nil 切片
	var sNil []int
	fmt.Println(sNil)

	//空切片
	sEmpty := make([]int, 0)
	fmt.Println(sEmpty)

	//空切片
	sEmpty = []int{}
	fmt.Println(sEmpty)

	//使用切片
	fmt.Println("使用切片")
	ss1 := []int{10, 20, 30, 40, 50}
	ss1[1] = 100
	fmt.Println(ss1)


}
