package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Printf("C\t%T \n", c)
	fmt.Printf("CR\t%T \n", cr)
	fmt.Printf("CS\t%T \n", cs)

	//convert general to specific channel
	fmt.Printf("C\t%T\n", (<-chan int)(c))
	fmt.Printf("C\t%T\n", (chan<- int)(c))
}
