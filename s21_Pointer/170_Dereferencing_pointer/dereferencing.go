package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%T %v\n", &x, &x)

	y := &x
	fmt.Printf("%T %v\n", y, y)

	//* show value = 42 instead of 0xc00001c098
	//* = Dereferencing
	fmt.Println(*y)

	//use *y to change value of x
	*y = 43
	fmt.Println(x)
}
