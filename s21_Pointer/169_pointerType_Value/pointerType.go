package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%T %v\n", &x, &x)

	s := "Jun"
	fmt.Printf("%T %v\n", &s, &s)
}
