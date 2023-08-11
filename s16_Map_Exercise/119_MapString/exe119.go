package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":       {"shaken", "not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	for key, mval := range m {
		fmt.Println("Key = ", key)
		for i, sval := range mval {
			fmt.Printf("Index %v = %v \n", i, sval)
		}
	}
}
