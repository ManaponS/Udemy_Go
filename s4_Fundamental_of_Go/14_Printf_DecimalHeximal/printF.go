package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	arr1 := []int{0, 1, 2, 3, 4, 5}
	for _, val := range arr1 {
		fmt.Printf("binary of %v : %b \n", val, val)
		fmt.Printf("hexadecimal of %v : %x \n", val, val)
		fmt.Println("-------------------------------------------")
	}
}
