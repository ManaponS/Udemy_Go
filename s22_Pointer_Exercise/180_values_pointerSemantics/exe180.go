package main

import "fmt"

type person struct {
	first string
}

func changeNameV(p person, s string) person {
	p.first = s
	return p
}
func changeNameP(p *person, s string) {
	p.first = s
}

func main() {
	p1 := person{
		first: "Leroys",
	}
	fmt.Println(p1)

	p1 = changeNameV(p1, "Jenkins")
	fmt.Println(p1)

	changeNameP(&p1, "Ramuh")
	fmt.Println(p1)

}
