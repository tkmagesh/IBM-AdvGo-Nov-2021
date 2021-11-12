package main

import (
	"fmt"
	"sync"
)

//share memory for communication
var invocationCount int

var wg sync.WaitGroup = sync.WaitGroup{}
var mutex sync.Mutex = sync.Mutex{}

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
	mutex.Lock()
	invocationCount++
	mutex.Unlock()
}
