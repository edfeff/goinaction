package main

import (
	"bytes"
	"fmt"
	"github.com/wpp/goinaction/chapter05/inserttype"
	"io"
	"net/http"
	"os"
)

//func init() {
//	if len(os.Args) != 2 {
//		fmt.Println("usage： ./main url")
//		os.Exit(-1)
//	}
//}

//go 的类型系统
func main() {
	//gotype.Show()
	//gomethod.Show()
	//gotype.InnerTypeShow()
	//testBuffer()
	//interfaces.Show()
	inserttype.Show()
}
func testCurl() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(resp)
		return
	}
	io.Copy(os.Stdout, resp.Body)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}

}
func testBuffer() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, "World")
	io.Copy(os.Stdout, &b)
}
