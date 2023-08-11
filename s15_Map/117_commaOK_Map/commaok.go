package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	delete(m, "a")

	//if print "a" the return value will be 0
	fmt.Println("a =", m["a"])
	v, ok := m["a"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key doesn't exist")
	}
}
