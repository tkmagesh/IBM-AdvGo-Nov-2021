/* Higher order functions */

package main

import (
	"fmt"
	"time"
)

type OperFunc func(int, int) int

func main() {
	/*
		fmt.Println(add(100, 200))
		fmt.Println(subtract(100, 200))
	*/
	/* fmt.Println(invokeAdd(100, 200))
	fmt.Println(invokeSubtract(100, 200)) */

	/*
		fmt.Println(invoke(add, 100, 200))
		fmt.Println(invoke(subtract, 100, 200))
	*/

	/* logAdd := logger(add)
	logSubtract := logger(subtract)
	fmt.Println(logAdd(100, 200))
	fmt.Println(logSubtract(100, 200)) */

	/*
		profAdd := profile(add)
		profSubtract := profile(subtract)
		fmt.Println(profAdd(100, 200))
		fmt.Println(profSubtract(100, 200))
	*/
	proflileLoggedAdd := profile(logger(add))
	proflileLoggedSubtract := profile(logger(subtract))
	fmt.Println(proflileLoggedAdd(100, 200))
	fmt.Println(proflileLoggedSubtract(100, 200))
}

func add(x, y int) int {
	time.Sleep(3 * time.Second)
	return x + y
}

/* func invokeAdd(x, y int) int {
	return add(x, y)
} */

func subtract(x, y int) int {
	time.Sleep(4 * time.Second)
	return x - y
}

/* func invokeSubtract(x, y int) int {
	return subtract(x, y)
} */

/*
func invoke(oper func(int, int) int, x, y int) int {
	fmt.Println("Before invocation")
	result := oper(x, y)
	fmt.Println("After invocation")
	return result
}
*/

/*
func logger(oper func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("After invocation")
		return result
	}
}
*/

func logger(oper OperFunc) OperFunc {
	var fn OperFunc = func(x, y int) int {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("After invocation")
		return result
	}
	return fn
}

func profile(oper OperFunc) OperFunc {
	var fn OperFunc = func(x, y int) int {
		start := time.Now()
		result := oper(x, y)
		end := time.Now().Sub(start) / time.Millisecond
		fmt.Printf("Operation took %d milliseconds\n", end)
		return result
	}
	return fn
}

/* Write a function that will profile the execution time of add & subtract */
/* combine the profile with logger */
