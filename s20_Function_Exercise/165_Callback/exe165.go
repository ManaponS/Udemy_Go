package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 4))
}

func square(n int) int {
	return n * n
}
func printSquare(f func(int) int, a int) string {
	return fmt.Sprintf("%v^2 = %v", a, f(a))
}

// func main() {
// 	fmt.Println(doMath(4, 8, multiply))
// 	fmt.Println(doMath(5, 3, mod))
// }
// func doMath(a int, b int, f func(int, int) int) int {
// 	return f(a, b)
// }

// func multiply(a, b int) int {
// 	return a * b
// }

// func mod(a, b int) int {
// 	return a % b
// }
