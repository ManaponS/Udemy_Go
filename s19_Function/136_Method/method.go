package main

import "fmt"

type person struct {
	name string
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func main() {
	p1 := person{
		name: "Jane",
	}
	p2 := person{
		name: "Res",
	}

	p1.speak()
	p2.speak()

}
