package main

import "fmt"

func returnFn() func() int {
	return func() int {
		return 42
	}
}

func main() {
	r := returnFn()
	fmt.Println(r())
}
