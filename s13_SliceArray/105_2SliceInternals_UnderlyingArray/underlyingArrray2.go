package main

import "fmt"

func main() {
	c := []int{0, 1, 2, 3, 4, 5}

	d := make([]int, len(c))
	copy(d, c)
	fmt.Printf("C : %v \n", c)
	fmt.Printf("D : %v \n", d)
	fmt.Println("--------------------")

	c[0] = 888
	fmt.Printf("C : %v \n", c)
	fmt.Printf("D : %v \n", d)
	fmt.Println("--------------------")
}
