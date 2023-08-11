package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type employee struct {
	person
	age  int
	role string
}

func main() {
	emp1 := employee{
		person: person{
			firstName: "John",
			lastName:  "Soda",
			age:       28,
		},
		age:  35,
		role: "Bartender",
	}

	fmt.Printf("%T \t %v \n", emp1, emp1)

	//will give age = 35
	fmt.Println(emp1.firstName, emp1.lastName,
		emp1.age, emp1.role)

	fmt.Println(emp1.person.age)
}
