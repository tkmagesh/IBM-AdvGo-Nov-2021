package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int) {
	defer close(ch)

	/* done := func() chan bool {
		doneCh := make(chan bool)
		go func() {
			time.Sleep(20 * time.Second)
			doneCh <- true
		}()
		return doneCh
	}() */

	done := time.After(20 * time.Second)

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-done:
			fmt.Println("timeout reached!")
			return
		}
	}
}
