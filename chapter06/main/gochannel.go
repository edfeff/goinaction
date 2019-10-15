package main

import "fmt"

//通道
//使用通道来发送和接受共享的资源，在goroutine之间做同步
func TestGoChannel() {
	fmt.Println("创建通道")

	//int 的通道,无缓冲区，不能保存值，要求发送和接受方必须同时准备才可以用
	//两个通过的接受和发送都是同步的。
	ch1 := make(chan int)
	fmt.Println(ch1)

	//有缓冲的，缓冲区大小
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 120)

	//向通道发送值
	fmt.Println("向通道发送值")
	//ch1 <- 1
	ch2 <- "Hello"
	ch3 <- "World"

	//接收值
	//v1 := <-ch1
	v2 := <-ch2
	v3 := <-ch3
	fmt.Println(v2, v3)
}
