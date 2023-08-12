package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(t ...int) int {
	fmt.Printf("%T \t %v\n", t, t)
	var sum int
	for _, v := range t {
		sum += v
	}
	return sum
}
