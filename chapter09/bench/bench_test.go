package bench

import (
	"fmt"
	"testing"
)

//基准测试
//  可以看到第一个基准测试函数，名为BenchmarkSprintf。
//基准测试函数必须以Benchmark 开头，接受一个指向testing.B 类型的指针作为唯一参数。
//为了让基准测试框架能准确测试性能，它必须在一段时间内反复运行这段代码，所以这里使用了
//for 循环，
//go test -v -run="none" -bench="BenchmarkSprintf"
//在这次 go test 调用里，我们给-run 选项传递了字符串"none"，来保证在运行制订的基
//准测试函数之前没有单元测试会被运行。这两个选项都可以接受正则表达式，来决定需要运行哪
//些测试
func BenchmarkSprintf(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", n)
	}
}

//go test -v -run="none" -bench="BenchmarkSprintf"
//goos: windows
//goarch: amd64
//pkg: github.com/wpp/goinaction/chapter09/bench
//BenchmarkSprintf-4      20000000               108 ns/op
//PASS
//ok      github.com/wpp/goinaction/chapter09/bench       2.593s

//可以使用另一个名为-benchtime 的
//选项来更改测试执行的最短时间
