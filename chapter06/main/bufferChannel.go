package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var b3 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func TestBufferChannel() {
	//有缓冲区的通道
	tasks := make(chan string, taskLoad)
	b3.Add(numberGoroutines)
	//启动goroutine来处理工作
	//consumer
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	//producer
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task %d", post)
	}
	close(tasks)
	b3.Wait()
}

func worker(task chan string, wo int) {
	defer b3.Done()
	for {
		t, ok := <-task
		if !ok {
			fmt.Printf("worker%d over\n", wo)
			return
		}
		fmt.Printf("worker%d: %s\n", wo, t)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("worker%d,complete %s \n", wo, t)
	}
}
