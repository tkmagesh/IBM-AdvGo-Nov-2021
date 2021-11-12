package main

import (
	"fmt"
	"sync"
)

var result = []int{}
var mutex sync.Mutex = sync.Mutex{}
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go add(i*10, i*2*10)
	}
	wg.Wait()
	fmt.Println(result, len(result))
}

func add(x, y int) {
	r := x + y
	mutex.Lock()
	{
		result = append(result, r)
	}
	mutex.Unlock()
	wg.Done()
}

/*
	Perform the add operation concurrently.
	Accumulate the results in a slice
*/
