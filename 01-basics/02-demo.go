package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(no int) {
			fmt.Println("i = ", no)
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}
