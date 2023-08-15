package main

import (
	"encoding/json"
	"fmt"
)

// first letter must be Uppercase to make it able to import
// if not it will retrurn [{},{}]
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Eric",
		Last:  "Blunc",
		Age:   45,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
