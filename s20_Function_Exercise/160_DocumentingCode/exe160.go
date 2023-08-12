package main

import "fmt"

func main() {
	fmt.Println(paradise("Hawaii"))
}

// Test Doc
func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}

// Add calculates the sum of two integers.
func Add(x, y int) int {
	return x + y
}
