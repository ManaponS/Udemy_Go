package main

import "fmt"

//create interface to hold type constrain instead of directly add it to T
type myNumbers interface {
	int | float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(85, 15))
	fmt.Println(addT(35.4, 64.7))
}
