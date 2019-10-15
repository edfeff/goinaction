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

var noBufferWg sync.WaitGroup

func TestNoBufferChannel() {
	//无缓冲通道
	c := make(chan int)

	noBufferWg.Add(2)

	go noBufferChannelPlayer("A", c)
	go noBufferChannelPlayer("B", c)

	//发送一个数字,启动数据传送
	c <- 1

	noBufferWg.Wait()
	fmt.Println("finish")
}

func noBufferChannelPlayer(name string, c chan int) {
	defer noBufferWg.Done()
	for {
		//第二个参数表示通道的状态，关闭时为false
		ball, ok := <-c
		if !ok {
			fmt.Println(name, "win")
			return
		}
		//随机数决定
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println(name, "missed")
			//关闭通道
			close(c)
			return
		}
		fmt.Println(name, "play", ball)
		ball++
		c <- ball
	}
}
