package main

import (
	"github.com/wpp/goinaction/chapter09/handlers"
	"log"
	"net/http"
)

func main() {
	testHandlerRoutes()
}

//测试端点
func testHandlerRoutes() {
	handlers.Routes()
	log.Println("listening 4000")
	http.ListenAndServe("localhost:4000", nil)
}
