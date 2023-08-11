package main

import "fmt"

func main() {
	arr := [5]int{}

	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	for i, v := range arr {
		fmt.Printf("Index %v = %v type %T\n", i, v, v)
	}
}
