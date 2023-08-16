package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//if not use close(c) the range will wait for more val this will cause deadlock
	//range is also block
	//recieve
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to Exit")
}
