package main

import "fmt"

func intDelta(n *int) {
	*n += 1
}

func modifyInt(n int) {
	n = 999
}

func sliceDelta(s []int) {
	s[0] = 99
}

func mapDelta(mp map[string]int, s string) {
	mp[s] = 95
}

type Person struct {
	Name string
	Age  int
}

func modifyPerson(p Person) {
	p.Age = 30
}

func modifySlice(s []int) {
	s[0] = 100
	s = append(s, 200)
}

func main() {
	a := 42
	fmt.Println(a)

	intDelta(&a)
	fmt.Println(a)

	modifyInt(a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 15
	fmt.Println(m)
	mapDelta(m, "James")
	fmt.Println(m)

	//Mutability
	person := Person{Name: "Alice", Age: 25}
	modifyPerson(person)
	fmt.Println(person.Age)
	numbers := []int{1, 2, 3}
	modifySlice(numbers)
	fmt.Println(numbers)
}
