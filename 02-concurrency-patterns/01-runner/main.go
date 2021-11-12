package main

import (
	"fmt"
	"runner-demo/runner"
	"time"
)

func main() {
	timeout := 50 * time.Second
	r := runner.New(timeout)
	r.Add(createTask(1))
	r.Add(createTask(2))
	r.Add(createTask(3))
	r.Add(createTask(4))
	//r.Add(createTask(5))
	if er := r.Start(); er != nil {
		switch er {
		case runner.ErrTimeout:
			fmt.Println("Terminating due to - timeout")
		case runner.ErrInterrupt:
			fmt.Println("Terminating due to - interrupt")
		}
	}
	fmt.Println("Processor ended")
}

func createTask(id int) func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id*10) * time.Second)
	}
}
