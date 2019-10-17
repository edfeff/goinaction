package main

import "log"

func init() {
	//日志前缀
	log.SetPrefix("TRACE:")
	//日志格式 日期 时间 文件代码行号
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// Ldate 日期 2019/10/16
	//Lmicroseconds 毫秒 21:28:02.495614
	//Llongfile 文件与行号  /home/wpp/go/src/github.com/wpp/goinaction/chapter08/golog1/main/main.go:12
	//log.Lshortfile 文件短写 main.go:16  (会覆盖Llongfile)
	//LstdFlags 标准记录器 等价 	log.Ldate|log.Ltime  2019/10/16 21:32:37

}
func main() {
	//日志
	//标准日志记录器
	log.Println("message")

	//Fatalln 会输出信息,并结束程序(调用 os.Exit(1))
	log.Fatalln("fatal message")

	//在调用Println之后,再调用panic
	log.Panicln("panic message")

	//输出
	// TRACE:2019/10/16 21:32:37 /home/wpp/go/src/github.com/wpp/goinaction/chapter08/golog1/main/main.go:19: message
	//TRACE:2019/10/16 21:32:37 /home/wpp/go/src/github.com/wpp/goinaction/chapter08/golog1/main/main.go:22: fatal message
}
