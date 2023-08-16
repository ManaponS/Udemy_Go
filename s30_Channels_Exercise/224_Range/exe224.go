package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("I ", i)
			c <- i
		}
		fmt.Println("Goroutine", runtime.NumGoroutine())

		close(c)
	}()
	return c
}
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
