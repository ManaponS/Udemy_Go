package main

import (
	"fmt"

	"github.com/ManaponS/Udemy_Go/s36_Testing_Benchmark_Exercise/254_/quote"
	"github.com/ManaponS/Udemy_Go/s36_Testing_Benchmark_Exercise/254_/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
