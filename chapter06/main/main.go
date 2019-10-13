package main

import "runtime"

func main() {
	//一个逻辑处理器,对于短时间的函数，一般不会发生函数切换，即函数会依次执行完毕再切换
	//TestGoRoutine(1)

	//但是对于耗时长的函数，有可能函数没有执行完毕，就发生了调度器的切换
	//TestLongTimeProgram(1)

	//TestDefer()

	//多个逻辑处理器，各个函数之间的顺序无法保证
	TestGoRoutine(runtime.NumCPU()) //大小写字符交错出现

}
