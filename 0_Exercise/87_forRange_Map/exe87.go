package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    42,
		"Moriarty": 32,
	}
	for index, val := range m {
		fmt.Printf("key: %v value: %v \n", index, val)
	}
}
