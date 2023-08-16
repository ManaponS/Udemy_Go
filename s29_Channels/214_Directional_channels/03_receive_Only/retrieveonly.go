package main

import "fmt"

func main() {
	c := make(<-chan int, 1)
	//invalid operation: cannot send to receive-only channel c (variable of type <-chan int)
	//c <- 42
	fmt.Println(<-c)

	fmt.Println("----------------")

	fmt.Printf("%T \n", c)
}
