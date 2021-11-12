package main

import (
	"fmt"
	"time"
	"worker-demo/worker"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"magesh",
	"ramesh",
	"suresh",
	"rajesh",
	"ganesh",
}

type namePrinter struct {
	name string
}

func (np *namePrinter) Task() {
	fmt.Println("Name Printer - Name : ", np.name)
	time.Sleep(2 * time.Second)
}

func main() {
	p := worker.New(5)
	for i := 0; i < 2; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			p.Run(&np)
		}
	}
	fmt.Println("All tasks are assigned")
	p.Shutdown()
}
