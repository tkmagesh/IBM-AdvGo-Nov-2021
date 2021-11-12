package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pooling-demo/pool"
	"sync"
	"sync/atomic"
	"time"
)

//Resource to be pooled
type DBConnection struct {
	ID int32
}

func (dbConn *DBConnection) Close() error {
	fmt.Println("Close : Connection : ", dbConn.ID)
	return nil
}

var idCounter int32

//resource factory
func DBConnectionFactory() (io.Closer, error) {
	fmt.Println("DBConnectionFactory : Creating a new DBConnection")
	atomic.AddInt32(&idCounter, 1)
	dbConn := &DBConnection{
		ID: idCounter,
	}
	return dbConn, nil
}

func main() {
	clientCount := 10
	wg := sync.WaitGroup{}
	wg.Add(clientCount)

	p, err := pool.New(DBConnectionFactory, 3)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Round 1 :")
	for client := 0; client < 5; client++ {
		go func(idx int) {
			doWork(idx, p)
			wg.Done()
		}(client)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("Round 2 :")
	for client := 0; client < 5; client++ {
		go func(idx int) {
			doWork(idx, p)
			wg.Done()
		}(client)
	}
	wg.Wait()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)

	fmt.Println("Worker : ", id, " : Acquired : ", conn.(*DBConnection).ID)
	//do some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Worker : ", id, " : Done : ", conn.(*DBConnection).ID)
}
