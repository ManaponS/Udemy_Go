package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    42,
		"Moriarty": 32,
	}

	age1 := m["James"]
	fmt.Println("Age of James", age1)

	age2 := m["Q"]
	fmt.Println("Age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q and the value is: ", v)
	}
}
