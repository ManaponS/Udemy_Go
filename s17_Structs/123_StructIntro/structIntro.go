package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "Anon",
		lastName:  "ymous",
		age:       28,
	}

	fmt.Printf("%T \t %v \n", p1, p1)

	fmt.Println(p1.firstName, p1.lastName, p1.age)

}
