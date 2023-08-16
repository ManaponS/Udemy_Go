package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	//retrive val from channel to print
	fmt.Println(<-c)
}
