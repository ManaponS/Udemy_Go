package main

import (
	"fmt"
)

// fix the deadlock
func main() {
	//buffer channel
	//c := make(chan int, 1)
	c := make(chan int)

	//func literal
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
