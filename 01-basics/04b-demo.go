package main

import (
	"fmt"
	"sync"
)

//share memory for communication
type InvocationCounter struct {
	invocationCount int
	sync.Mutex
}

func (i *InvocationCounter) Increment() {
	i.Lock()
	{
		i.invocationCount++
	}
	i.Unlock()
}

var wg sync.WaitGroup = sync.WaitGroup{}
var mutex sync.Mutex = sync.Mutex{}
var ic InvocationCounter = InvocationCounter{}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go doSomething()
	}
	wg.Wait()
	fmt.Println(ic.invocationCount)
}

func doSomething() {
	defer wg.Done()
	ic.Increment()
}
