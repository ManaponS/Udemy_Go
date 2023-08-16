package main

import "fmt"

func main() {
	c := make(chan<- int, 1)
	c <- 42
	//can't receive
	//invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	//fmt.Println(<-c)

	fmt.Println("----------------")

	fmt.Printf("%T \n", c)
}
