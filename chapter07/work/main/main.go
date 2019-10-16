package main

import (
	"github.com/wpp/goinaction/chapter07/work"
	"log"
	"sync"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

//实现worker接口
func (n *namePrinter) Task() {
	log.Println(n.name)
	//time.Sleep(time.Second)
}
func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
