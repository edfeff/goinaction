package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//没有缓冲区的通道

var bb sync.WaitGroup

func TestNoBufferChannel2() {
	//无缓冲通道
	c := make(chan int)
	//
	bb.Add(1)
	go run(c)
	//发送一个数字,启动数据传送
	c <- 1
	bb.Wait()
	fmt.Println("finish")
}

func run(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Println("next runner is ", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Println("now is ", newRunner)
		go run(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Println("Run over ", runner)
		bb.Done()
		return
	}
	fmt.Println("run ex with ", runner, newRunner)
	baton <- newRunner
}
