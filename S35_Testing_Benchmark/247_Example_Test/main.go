package main

import (
	"fmt"

	"github.com/ManaponS/Udemy_Go/S35_Testing_Benchmark/247_Example_Test/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 4, 6, 8, 10))
}
