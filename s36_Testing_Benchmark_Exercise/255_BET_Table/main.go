package main

import (
	"fmt"

	"github.com/ManaponS/Udemy_Go/s36_Testing_Benchmark_Exercise/255_BET_Table/mymath"
)

func main() {
	xxi := gen()
	fmt.Println(xxi)
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
