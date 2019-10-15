package main

import (
	"fmt"
	"runtime"
	"sync"
)

//互斥锁
var (
	c1 int
	g2 sync.WaitGroup
	//定义一段互斥代码
	mutex sync.Mutex
)

func TestMutualExclusion() {
	g2.Add(2)
	go doWork2(1)
	go doWork2(2)
	g2.Wait()
	fmt.Println("final", c1)
}
func doWork2(n int) {
	defer g2.Done()
	for i := 0; i < 2; i++ {
		//加锁，类似java的Lock
		mutex.Lock()
		{
			v := c1
			runtime.Gosched()
			v++
			c1 = v
		}
		//解锁
		mutex.Unlock()
	}

}
