package main

import "fmt"

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Anon",
		lastName:  "ymous",
		age:       28,
	}

	fmt.Printf("%T \n", p1)

	fmt.Println(p1.firstName, p1.lastName, p1.age)

}
