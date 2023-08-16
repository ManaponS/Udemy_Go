package main

import "fmt"

func main() {
	c := make(chan int)

	go put(c)
	pull(c)

	fmt.Println("Exit")
}
func put(c chan int) <-chan int {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	return c
}
func pull(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
