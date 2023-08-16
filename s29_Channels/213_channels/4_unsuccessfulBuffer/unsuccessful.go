package main

import "fmt"

func main() {
	c := make(chan int, 1)
	//can have only one buffer
	c <- 42
	//pull off one value or make size to 2 will make this work
	//fmt.Println(<-c)
	c <- 43
	fmt.Println(<-c)

}
