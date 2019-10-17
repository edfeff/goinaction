package main

import (
	"net/http"
	"testing"
)

//go 的test 只会测试以 _test结尾的文件
//。一个测试函数
//必须是公开的函数，并且以Test 单词开头。不但函数名字要以Test 开头，而且函数的签名必
//须接收一个指向testing.T 类型的指针，并且不返回任何值。
//调用命令 go test -v 即可测试

//如果需要报告测试失败，但是并不想停止当前测试函数的执行，可以使用t.Error
//t.Fatal 方法不但报告这个单元测试已经失败，而且会向测试输出写一些消息，而后立刻停止这个测试函数的执行
//如果测试函数执行时没有调用过t.Fatal 或者t.Error 方法，就会认为测试通过了
func Test1(t *testing.T) {
	t.Log("test1 1111")
	t.Error("Err Test1")
	t.Log("test1 2222")
	//=== RUN   Test1
	//--- FAIL: Test1 (0.00s)
	//    list01_test.go:18: test1 1111
	//    list01_test.go:19: Err Test1
	//    list01_test.go:20: test1 2222
}

func Test2(t *testing.T) {
	t.Fatal("Test2")
	t.Log("Test2 222")
	t.Log("Test2 333")
	//=== RUN   Test2
	//--- FAIL: Test2 (0.00s)
	//    list01_test.go:24: Test2
}
func Test3(t *testing.T) {
	t.Log("Test3")
}
func Test4(t *testing.T) {
	t.Log("Test4")
}
func Test5(t *testing.T) {
	t.Log("Test5")
}
func Test6(t *testing.T) {
	t.Log("Test6")
}

func NoTestDownload(t *testing.T) {
	const checkMark = "\u2713"
	const ballotX = "\u2713"

	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200
	t.Log("11")
	{
		t.Logf("test: %s , %d", url, statusCode)
		resp, err := http.Get(url)
		if err != nil {
			t.Fatal("fail:", err)
		}
		t.Log("222")
		defer resp.Body.Close()
		if resp.StatusCode == statusCode {
			t.Logf("%d: %v", statusCode, checkMark)
		} else {
			t.Errorf("%d %v %v", statusCode, ballotX, resp.StatusCode)
		}
	}
}

func TestMockServerDownload(t *testing.T) {
	const checkMark = "\u2713"
	const ballotX = "\u2713"

	server := mockServer()
	defer server.Close()

	url := server.URL
	statusCode := 200
	t.Log("11")
	{
		t.Logf("test: %s , %d", url, statusCode)
		resp, err := http.Get(url)
		if err != nil {
			t.Fatal("fail:", err)
		}
		t.Log("222")
		defer resp.Body.Close()
		if resp.StatusCode == statusCode {
			t.Logf("%d: %v", statusCode, checkMark)
		} else {
			t.Errorf("%d %v %v", statusCode, ballotX, resp.StatusCode)
		}
	}
}
