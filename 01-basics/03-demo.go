package main

import "fmt"

func main() {
	result := add(100, 200)
	fmt.Println(result)
}

func add(x, y int) int {
	return x + y
}
