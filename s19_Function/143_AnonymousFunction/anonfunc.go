package main

import "fmt"

func main() {
	foo()

	// an anonymous function
	// func(p parameter(s)) (r return(s)) { code }
	func() {
		fmt.Println("Anonymous Function")
	}()

	//Anonymous Function with Parameter
	func(s string) {
		fmt.Println("Anonymous name = ", s)
	}("Joe")

	//Return Anonymous Function
	r := func() string {
		return "Return Anonymous"
	}()
	fmt.Println(r)

}

func foo() {
	fmt.Println("Foo")
}
