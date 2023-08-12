package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(xi...))
}

func sum(t ...int) int {
	fmt.Printf("%T \t %v\n", t, t)
	var sum int
	for _, v := range t {
		sum += v
	}
	return sum
}
