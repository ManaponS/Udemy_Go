package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)
	fmt.Println("Random number is ", x)

	switch {
	case x <= 100:
		fmt.Println("X between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("X between 101 and 200")
	case x > 201 && x <= 250:
		fmt.Println("X between 201 and 250")
	default:
		fmt.Println("It more than 250")
	}
}
