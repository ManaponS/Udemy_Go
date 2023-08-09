package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count int
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v total count %v X = 3 \n", i, count)
			count++
		}
	}

}
