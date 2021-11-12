package main

import "fmt"

func main() {
	result := []int{}
	for i := 1; i <= 100; i++ {
		r := add(i*10, i*2*10)
	}
	fmt.Println(result, len(result))
}

func add(x, y int) int {
	return x + y
}

/*
	Perform the add operation concurrently.
	Accumulate the results in a slice
*/
