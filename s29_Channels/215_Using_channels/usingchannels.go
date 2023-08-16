package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//send
	go send(c)

	//receive
	//go receive(c) will not work bc program will exit before recieve c
	//using receive to block until receive the val from send
	receive(c)

	fmt.Println("About to Exit")
}

// send
func send(c chan<- int) {
	c <- 42
}

// receive
func receive(c <-chan int) {
	fmt.Println(<-c)
}
