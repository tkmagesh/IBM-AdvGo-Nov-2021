package main

import (
	"fmt"
	"sync"
)

//using channel
/*
func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
*/

var result int
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int) {
	result = x + y
	wg.Done()
}
