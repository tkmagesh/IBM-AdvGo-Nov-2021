package main

import "fmt"

func main() {
	var increment = getCounter()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getCounter() func() int {
	var no int
	var increment = func() int {
		no++
		return no
	}
	return increment
}
