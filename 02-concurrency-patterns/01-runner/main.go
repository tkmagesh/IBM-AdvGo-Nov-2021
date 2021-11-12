package main

import (
	"fmt"
	"runner-demo/runner"
	"time"
)

func main() {
	timeout := 20 * time.Second
	r := runner.New(timeout)
	r.Add(createTask())
	r.Add(createTask())
	r.Add(createTask())
	r.Add(createTask())
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

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
