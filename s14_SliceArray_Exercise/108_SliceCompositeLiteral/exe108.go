package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range slice {
		fmt.Printf("Index %v = %v type %T\n", i, v, v)
	}
}
