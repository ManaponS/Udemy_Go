package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName:   "Ryan",
		lastName:    "McCree",
		favIceCream: []string{"Coffee", "Late", "Mocca"},
	}
	p2 := person{
		firstName:   "Robert",
		lastName:    "Smith",
		favIceCream: []string{"Lime", "Orange", "Strawberry"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, mVal := range m {
		fmt.Println(mVal)
		for _, lVal := range mVal.favIceCream {
			fmt.Println(mVal.firstName, mVal.lastName, lVal)
		}
	}
}
