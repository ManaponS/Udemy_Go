package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anon X")
	}()

	func(s string) {
		fmt.Println("Recieve :", s)
	}("String")

	x := func() string {
		return fmt.Sprintln("Return X")
	}()
	fmt.Println(x)
}
