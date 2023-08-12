package main

import "fmt"

func main() {

	fmt.Println(foo())
	fmt.Printf("%T \n", foo)

	y := bar()
	fmt.Println(y())
	fmt.Printf("%T \n", bar)
	fmt.Printf("%T \n", y)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
