package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}

	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}
func foo(v ...int) int {
	var sum int
	for _, v := range v {
		sum += v
	}
	return sum
}

func bar(v []int) int {
	var sum int
	for _, v := range v {
		sum += v
	}
	return sum
}
