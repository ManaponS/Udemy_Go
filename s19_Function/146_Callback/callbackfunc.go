package main

import "fmt"

func main() {
	//callback = pass function in as an argument of another function
	fmt.Printf("%T \n", add)
	fmt.Printf("%T \n", minus)
	fmt.Printf("%T \n", doMath)

	x := doMath(1, 2, add)
	fmt.Println(x)

	y := doMath(2, 1, minus)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
