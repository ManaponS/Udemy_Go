package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v i'm now %v\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Som",
		age:   45,
	}
	p1.speak()
}
