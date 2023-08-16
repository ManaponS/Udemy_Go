package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("Exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

//select will look which channel that it can pull val off then leave the select statement
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even Channel:", v)
		case v := <-o:
			fmt.Println("Odd Channel:", v)
		case v := <-q:
			fmt.Println("Quit Channel:", v)
			return
		}
	}
}
