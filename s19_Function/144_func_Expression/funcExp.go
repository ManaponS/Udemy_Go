package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Anonymous Function")
	}

	y := func(s string) {
		fmt.Println("Anonymous name = ", s)
	}

	x()
	y("Kazoo")

}
