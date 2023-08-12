package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f()) //f = 5
	fmt.Println("-------------------------------")
	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g()) //g = 5
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
