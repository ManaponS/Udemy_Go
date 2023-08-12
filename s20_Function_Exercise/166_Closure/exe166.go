package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(a float64) func() float64 {
	var b float64
	return func() float64 {
		b++
		fmt.Printf("%v pow %b = ", a, b)
		return (math.Pow(a, b))
	}
}
