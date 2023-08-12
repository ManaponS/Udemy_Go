package main

import "fmt"

type person struct {
	name string
}

type man struct {
	person
	mustache bool
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func (m man) speak() {
	fmt.Println("A man name", m.name)
}

type human interface {
	speak()
}

func saySMT(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Jane",
	}
	p2 := person{
		name: "Res",
	}

	m1 := man{
		person: person{
			name: "Chad",
		},
		mustache: true,
	}

	// p1.speak()
	// p2.speak()

	// m1.speak()
	saySMT(p1)
	saySMT(p2)
	saySMT(m1)

}
