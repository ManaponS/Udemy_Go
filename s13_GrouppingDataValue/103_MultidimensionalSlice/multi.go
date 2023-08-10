package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}

	fmt.Println(a)
	fmt.Println(b)

	x := [][]string{a, b}
	fmt.Println(x)
}
