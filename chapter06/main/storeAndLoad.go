package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	g1       sync.WaitGroup
)

func TestStoreAndLoad() {
	g1.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown now")
	//
	atomic.StoreInt64(&shutdown, 1)
	g1.Wait()
}

func doWork(s string) {
	defer g1.Done()
	for {
		fmt.Printf("Doing %s work\n", s)
		time.Sleep(250 * time.Millisecond)
		//多个goroutine之间可见这个shutdown
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shut.. %s\n", s)
			break
		}
	}
}
