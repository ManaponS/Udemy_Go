package main

import "fmt"

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "I'm 008."}

	s := [][]string{s1, s2}
	fmt.Println(s)

	for i1, v1 := range s {
		fmt.Println(i1, v1)
		for i2, v2 := range v1 {
			fmt.Println(i2, v2)
		}
	}
}
