package main

import "fmt"

var a int

func main() {
	fmt.Println(a)

	b := 1
	fmt.Println(b)

	c, d := 1+2i, "GGS"
	fmt.Println(c, d)

	var e float64 = 6.09
	fmt.Println(e)

	f, _ := true, false
	fmt.Println(f)
}
