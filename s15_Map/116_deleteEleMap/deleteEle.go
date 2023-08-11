package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	fmt.Println("---------------------------")

	delete(m, "a")
	fmt.Println(m)

	//if print "a" the return value will be 0
	fmt.Println("a =", m["a"])
}
