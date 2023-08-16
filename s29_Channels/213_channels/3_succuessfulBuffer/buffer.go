package main

import "fmt"

func main() {
	//buffer channel let some val sit in regardless of having smt to get it
	c := make(chan int, 1)
	//c can have 1 buffer <- 42
	c <- 42
	fmt.Println(<-c)
}
