package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"dJames", 32}
	p2 := person{"cMoneypenny", 27}
	p3 := person{"bQ", 64}
	p4 := person{"aM", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].first < people[j].first
	})
	fmt.Println(people)
}
