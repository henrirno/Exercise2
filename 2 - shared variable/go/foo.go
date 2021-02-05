// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int, readyToRun <-chan bool) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case msg1 := <-add:
			number = number + msg1
		case msg2 := <-sub:
			number = number - msg2
		case <-readyToRun:
			read <- number
		}

	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000+2; j++ {
		sub <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	read := make(chan int)
	finished := make(chan bool)
	add := make(chan int)
	sub := make(chan int)
	readyToRun := make(chan bool)

	// TODO: Spawn the required goroutines
	go incrementer(add, finished)
	go decrementer(sub, finished)
	go number_server(add, sub, read, readyToRun)
	// TODO: block on finished from both "worker" goroutines
	<-finished
	<-finished
	readyToRun <- true

	fmt.Println("The magic number is:", <-read)
}
