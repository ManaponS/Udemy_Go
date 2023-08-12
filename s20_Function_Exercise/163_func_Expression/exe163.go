package main

import "fmt"

func main() {

	fmt.Println(x(9, 6))
}

var x = func(i1, i2 int) int {
	return i1 + i2
}
