package main

import (
	"fmt"
	"sync"
)

var result int
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int, 1)
	wg.Add(1)
	go add(100, 200, ch)
	//modify the sequence of the following 2 lines of code ( try with a non-buffered channel and try with a buffered channel)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result = x + y
	ch <- result
	wg.Done()
}
