package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

//输入输出
func main() {
	writerTest()
	readerTest()
	curlTest()
}

//接口定义
// type Writer interface {
//	Write(p []byte) (n int, err error)
//}
//
// type Reader interface {
//	Read(p []byte) (n int, err error)
//}
func writerTest() {
	//Buffer实现了Writer接口
	var b bytes.Buffer

	//直接写入
	b.Write([]byte("hello"))

	//利用fmt追加
	fmt.Fprintf(&b, " World")

	//输出到stdout中
	b.WriteTo(os.Stdout)
}

func readerTest() {

}
func curlTest() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("baidu.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//组合多个输出设备
	dest := io.MultiWriter(os.Stdout, f)

	v, err := io.Copy(dest, r.Body)
	fmt.Println(v)
	if err != nil {
		fmt.Println(err)
	}

}
