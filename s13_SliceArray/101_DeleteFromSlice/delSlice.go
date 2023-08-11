package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)
	fmt.Println("-------------")

	xi = append(xi[:1], xi[2:]...)
	fmt.Println(xi)
	fmt.Println("-------------")

}
