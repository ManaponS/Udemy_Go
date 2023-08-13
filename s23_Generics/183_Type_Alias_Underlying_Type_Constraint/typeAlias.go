package main

import "fmt"

//~ make all int and float include in interface
type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 42
	//fmt.Println(addT(n, 15))
	//myAlias does not satisfy myNumbers (possibly missing ~ for int in myNumbers)
	fmt.Println(addT(n, 15))
	fmt.Println(addT(35.4, 64.7))
}
