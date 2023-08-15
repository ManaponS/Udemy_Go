package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi I'm ", p.first, "I now", p.age)
}

type human interface {
	speak()
}

func saySMT(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Oleg",
		age:   54,
	}
	p2 := &person{
		first: "Desmond",
		age:   22,
	}
	//not work
	// saySMT(p1)
	saySMT(&p1)
	saySMT(p2)
}
