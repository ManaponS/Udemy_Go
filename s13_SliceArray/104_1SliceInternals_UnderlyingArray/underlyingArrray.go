package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	b := a

	fmt.Printf("A : %v \n", a)
	fmt.Printf("B : %v \n", b)
	fmt.Println("--------------------")

	a[0] = 999
	fmt.Printf("A : %v \n", a)
	fmt.Printf("B : %v \n", b)
	fmt.Println("--------------------")
	//a and b use same undetlying array

}
