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

	//使用切片创建切片
	ss2 := ss1[1:3]
	ss2[0] = 999
	ss2[1] = 999
	fmt.Println(ss1)
	fmt.Println(len(ss1))
	fmt.Println(ss2)
	fmt.Println(len(ss2))

	//切片扩容
	//长度和容量都是5
	ss3 := []int{1, 2, 3, 4, 5}
	//长度是2，容量是 5-1 =4
	ss4 := ss3[1:3]
	fmt.Println(ss4)
	//使用append会修改切片的尺寸
	ss4 = append(ss4, 6)
	ss4 = append(ss4, 7)
	ss4 = append(ss4, 8)
	ss4 = append(ss4, 9)
	ss4 = append(ss4, 10)
	ss4 = append(ss4, 11)
	ss4 = append(ss4, 12)
	ss4 = append(ss4, 13)

	//使用append时，也会修改关联的切片
	fmt.Println(ss4) //[2 3 6 7 8 9 10 11 12 13]
	fmt.Println(ss3) //[1 2 3 6 7]

	//限制切片的容量
	//5个元素
	ss5 := []string{"a", "b", "c", "d", "e"}
	//使用3个索引，分别是 [开始位置,结束位置) 最终位置
	//实际长度 = 结束位置-开始位置
	//实际容量 = 最终位置 - 开始位置
	//长度为1 容量为2的
	//这样会在修改时，发现新的切片和原有切片共用了开始位置和结束位置之间的内存，
	// 如果共享位置的值发生了修改，新切片和旧切片的影响都是相同的，
	//如果新切片的长度超出了指定的共享位置，那么就不再和就切片有联系
	ss6 := ss5[2:3:4]
	ss6[0] = "X"
	ss6 = append(ss6, "f")
	ss6 = append(ss6, "g")
	fmt.Println(ss6) //[X f g]
	fmt.Println(ss5) //[a b X f e]
	//所以，新旧切片出现了共享空间就会发生一些列诡异的问题，所以要避免这样使用
	//一般将新切片的结束位置等于最种位置，这样当使用append函数时，go会开辟一段新的内存地址

	ss7 := ss5[2:3:3]
	ss7 = append(ss7, "7")
	ss7 = append(ss7, "7")
	fmt.Println(ss5) //[a b X f e]
	fmt.Println(ss7) //[X 7 7]

	//切片追加
	ss8 := []int{1, 2}
	ss9 := []int{3, 4}
	ss10 := append(ss8, ss9...)
	fmt.Println(ss8)
	fmt.Println(ss9)
	fmt.Println(ss10)

	//迭代切片
	//range
	ss11 := []int{10, 20, 30, 40}
	for index, value := range ss11 {
		fmt.Println(index, value)
	}
	//不使用索引
	for _, value := range ss11 {
		fmt.Println(value)
	}
	//不使用值
	for index, _ := range ss11 {
		fmt.Println(index)
	}
	//range 提供的是副本
	//比较下内存地址
	for index, value := range ss11 {
		fmt.Printf("%X,%X \n", &value, &ss11[index])
	}
	//传统for循环,根精确的索引控制
	for index := 1; index < len(ss11); index += 2 {
		fmt.Println(ss11[index])
	}
	//切片的长度和容量
	//len
	//cap
	fmt.Println(len(ss11))
	fmt.Println(cap(ss11))

	//多维切片
	m1 := [][]int{{10}, {100, 200}}
	fmt.Println(m1)

	//组合切片的切片
	m1[0] = append(m1[0], 20)
	fmt.Println(m1)

	//在函数种使用切片参数
	f1()
}
func f1() {
	s := make([]int, 1e6)
	//函数间传递推荐使用切片，因为切片本身的结构就3部分，头指针，长度，容量，
	// 所以当作参数传递比较简单，比较节省空间
	s2 := f2(s)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
func f2(slice []int) []int {
	return slice
}
