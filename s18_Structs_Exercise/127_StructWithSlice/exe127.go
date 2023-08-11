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

	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.favIceCream {
		fmt.Println("P1 favorite ice cream flavor is : ", v)
	}

	for _, v := range p2.favIceCream {
		fmt.Println("P2 favorite ice cream flavor is : ", v)
	}

}
