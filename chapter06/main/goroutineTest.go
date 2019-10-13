package main

import (
	"fmt"
	"runtime"
	"sync"
)

var g sync.WaitGroup

func TestLongTimeProgram(num int) {
	//设置调度器的最大逻辑处理器个数为1
	runtime.GOMAXPROCS(num)
	g.Add(2)
	fmt.Println("start")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("wait..")
	g.Wait()
	fmt.Println("finished")
}

func printPrime(s string) {
	defer g.Done()
	//找素数
next:
	for i := 2; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", s, i)
	}
	fmt.Println("Completed", s)
}
func TestDefer() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
func TestGoRoutine(num int) {
	//分配一个逻辑处理器
	runtime.GOMAXPROCS(num)

	//等待goroutine 为2个时
	//用于计算goroutine的数量
	//类似java 的countDownLatch
	//Add 方法Add为等待的goroutine的个数
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine")
	//匿名函数
	go func() {
		//函数退出时，调用Done来通知main已经完成
		//类似java中的CountDownLatch
		//Done则类似countDown
		//defer关键字会改变代码执行的顺序
		//只有当前函数执行完毕后才会执行defer定义的代码
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	fmt.Println("Wait To Finish")
	wg.Wait()
	fmt.Println("\n terminating Program")
}
