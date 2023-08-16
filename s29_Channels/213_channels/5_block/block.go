package main

import (
	"fmt"
	"time"
)

func add(c chan int, a, b int) {
	result := a + b
	fmt.Println("c = ", result)
	c <- result
}

func multiply(c chan int) {
	result := <-c
	result *= 2

	fmt.Printf("Result is %d\n", result)
}

func main() {
	ch := make(chan int)
	a := 10
	b := 5

	go add(ch, a, b)
	go multiply(ch)

	time.Sleep(1 * time.Second)
}
