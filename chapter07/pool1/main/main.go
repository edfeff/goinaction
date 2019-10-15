package main

import (
	"github.com/wpp/goinaction/chapter07/pool1/mypool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close:Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	p, err := mypool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	for q := 0; q < maxGoroutines; q++ {
		go func(q int) {
			performQ(q, p)
			wg.Done()
		}(q)
	}
	wg.Wait()
	log.Println("ShutDown")
	p.Close()
}

func performQ(query int, pool *mypool.Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Println("QID:", query, conn.(*dbConnection).ID)
}
