package array

//go 的数组
import "fmt"

//数组
var arr [5]int

//初始化
func init() {
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
}

//Show 遍历并取值
func Show() {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//声明并赋值
	arr2 := [2]int{10, 20}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	//自动长度
	arr3 := [...]int{31, 32, 33, 34}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//个别元素赋值
	arr4 := [5]int{1: 10, 4: 44}
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}
	//指针数组
	arr5 := [5]*int{0: new(int), 1: new(int)}
	*arr5[0] = 10
	*arr5[1] = 10

	//访问指针数组的元素
	fmt.Println(*arr5[0])
	fmt.Println(*arr5[0])

	//数组之间的赋值
	var arr6 [5]string
	arr7 := [5]string{"a", "b", "c", "d", "e"}
	arr6 = arr7
	arr6[1] = "X"
	fmt.Println(arr6[1], arr7[1]) //X b 是深度复制

	//指针数组赋值
	var arr8 [3]*string
	arr9 := [3]*string{new(string), new(string), new(string)}
	*arr9[0] = "a"
	*arr9[1] = "b"
	*arr9[2] = "c"
	arr8 = arr9
	*arr8[0] = "X"
	fmt.Println(*arr8[0], *arr9[0]) //X X 指向同一片内存

	//多维数组
	var arr10 [4][2]int
	//单个赋值
	arr10[0][0] = 100
	arr11 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	//对外层赋值
	arr12 := [4][2]int{0: {10, 11}, 3: {40, 41}}
	//对内层赋值
	arr13 := [4][2]int{
		0: {0: 10, 1: 11},
		3: {0: 40, 1: 41}}
	//取出二维数组的值
	fmt.Println(arr10[0][0]) //0
	fmt.Println(arr11[0][1]) //11
	fmt.Println(arr12[3][0]) //40
	fmt.Println(arr13[3][1]) //41

	//多维数组赋值-也是深度复制
	var arr15 [4][2]int
	arr15 = arr13
	arr15[0][0] = 15
	fmt.Println(arr15[0][0], arr13[0][0])

	//一层赋值
	arr16 := [2][2]int{{16, 16}, {16, 16}}
	arr17 := [2][2]int{{17, 17}, {17, 17}}
	arr16[0] = arr17[0]
	fmt.Println()
	for i := 0; i < len(arr16); i++ {
		for j := 0; j < len(arr16[0]); j++ {
			fmt.Printf("%d ", arr16[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(arr16); i++ {
		for j := 0; j < len(arr16[0]); j++ {
			fmt.Printf("%d ", arr17[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	//向函数传递参数，传递数组指针，不能对数组进修改
	fooShow(arr)
	fooArray(arr)
	fooShow(arr)

	//向函数传递参数，传递数组指针,里面可以对数组进行修改
	fooPointer(&arr)
	fooShow(arr)

}

//fooArray 参数是深度复制的
func fooArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 999 + arr[i]
	}
}

//在函数件传递数组
func fooPointer(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 5555 + arr[i]
	}
}

func fooShow(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
