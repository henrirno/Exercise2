// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
        }
	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	//TODO: signal that the goroutine is finished
}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	//TODO: signal that the goroutine is finished
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	read := make(chan int)

	// TODO: Spawn the required goroutines

	// TODO: block on finished from both "worker" goroutines

	fmt.Println("The magic number is:", <-read)
}
