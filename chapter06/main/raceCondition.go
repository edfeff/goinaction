package main

import (
	"fmt"
	"runtime"
	"sync"
)

//并发的竞争条件
var (
	counter int
	wg      sync.WaitGroup
)

func TestRaceCondition() {
	wg.Add(2)
	go intCounter(1)
	go intCounter(2)
	wg.Wait()
	fmt.Println("Final :", counter)
	//结果是2，正确应该是4，
	// 因为goroutine切换时，值进行了覆盖
}
func intCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		v := counter
		//当前routine退出，回到队列中,模拟激烈的并发
		runtime.Gosched()
		v++
		counter = v
	}
}
