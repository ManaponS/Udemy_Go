package main

import "fmt"

//add T as Type constrain that specify the type that want to take
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(85, 15))
	fmt.Println(addT(35.4, 64.7))
}
