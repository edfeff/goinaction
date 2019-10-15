package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter2 int64
	wg2      sync.WaitGroup
)

func TestLock() {
	wg2.Add(2)
	go intCounter2(1)
	go intCounter2(2)
	wg2.Wait()
	fmt.Println("Final :", counter2)
	//结果是2，正确应该是4，
	// 因为goroutine切换时，值进行了覆盖
}
func intCounter2(id int) {
	defer wg2.Done()
	for count := 0; count < 2; count++ {
		//使用原子类安全的增加
		atomic.AddInt64(&counter2, 1)
		//当前routine退出，回到队列中,模拟激烈的并发
		runtime.Gosched()
	}
}
