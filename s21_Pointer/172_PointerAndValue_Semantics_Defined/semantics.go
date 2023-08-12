package main

import "fmt"

//Value Semantics
func addOneV(v int) int {
	return v + 1
}

//Pointer Semantics
func addOneP(p *int) {
	*p += 1
}

func main() {
	//Value Semantics
	a := 1
	fmt.Println(a) //1
	/*
		when the actual data of a variable is passed to a function or assigned to another variable. This means that the new variable or function parameter gets a completely independent copy of the data - EVERYTHING IN GO IS 'PASS BY VALUE'.
	*/
	fmt.Println(addOneV(a)) //2
	fmt.Println(addOneV(a)) //2
	fmt.Println(a)          //1
	fmt.Println("-------------------")

	//Pointer Semantics
	b := 5
	fmt.Println(b) //5
	addOneP(&b)
	fmt.Println(b) //6
	addOneP(&b)
	fmt.Println(b) //7
}
