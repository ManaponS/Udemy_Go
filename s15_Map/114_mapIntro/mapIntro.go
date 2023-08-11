package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println("Value of A =", mp["A"])

	fmt.Printf("%T \t %v \n", mp, mp)

	makeMap := make(map[string]int)
	makeMap["D"] = 4
	makeMap["E"] = 5
	makeMap["F"] = 6

	fmt.Println(len(makeMap))

	fmt.Println("Value of A =", makeMap["D"])

	fmt.Printf("%T \t %v \n", makeMap, makeMap)
}
