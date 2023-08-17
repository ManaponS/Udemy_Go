package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("2 + ...5 =", mySum(2, 3, 4, 5))
	fmt.Println("14 + 62 =", mySum(43, 62))

}
func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
