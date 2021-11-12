package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//share memory for communication
var invocationCount int64

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go doSomething()
	}
	wg.Wait()
	fmt.Println(invocationCount)
}

func doSomething() {
	defer wg.Done()
	atomic.AddInt64(&invocationCount, 1)
}
