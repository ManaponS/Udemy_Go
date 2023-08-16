package main

import (
	"fmt"
)

func main() {
	q := make(chan bool)

	c := gen(q)
	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- bool) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- false
		close(c)
	}()
	return c
}

func receive(c <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
