package main

import (
	"github.com/wpp/goinaction/chapter07/runner/runnerpackage"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("start...")
	r := runnerpackage.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runnerpackage.ErrTimeout:
			log.Println("stop by timeout")
			os.Exit(1)
		case runnerpackage.ErrInterrupt:
			log.Println("stop by interrupt")
			os.Exit(2)
		}
	}
	log.Println("end")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("T-%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
