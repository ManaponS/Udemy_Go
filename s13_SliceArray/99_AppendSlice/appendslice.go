package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("-------------")

	xi = append(xi, 45, 46, 47)
	fmt.Println(xi)
	fmt.Println("-------------")

	xi = append(xi, 41)
	fmt.Println(xi)
	fmt.Println("-------------")

}
