package main

import (
	"fmt"

	"github.com/ManaponS/Udemy_Go/s36_Testing_Benchmark_Exercise/253_BET/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
