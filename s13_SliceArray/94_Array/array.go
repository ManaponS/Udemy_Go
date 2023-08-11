package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{"a", "b", "c"}
	fmt.Println(b)

	var c [2]int
	c[0] = 1
	c[1] = 2
	fmt.Println(c)

	fmt.Printf("Type of A : %T \n", a)
	fmt.Printf("Type of B : %T \n", b)
	fmt.Printf("Type of C : %T \n", c)

	fmt.Println(len(a))

}
