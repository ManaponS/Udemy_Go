package main

import "fmt"

func main() {
	a()
	b()
	c()
}

// The behavior of defer statements is straightforward and predictable. There are three simple rules:
// 1.A deferred function’s arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	//In this example, the expression “i” is evaluated when the Println call is deferred. The deferred call will print “0” after the function returns.
	defer fmt.Println(i)
	i++
}

//2.Deferred function calls are executed in Last In First Out order after the surrounding function returns.

func b() {
	for i := 0; i < 4; i++ {
		//This function prints “3210”:
		defer fmt.Print(i)
	}
}

//3.Deferred functions may read and assign to the returning function’s named return values.
func c() (i int) {
	//In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:
	defer func() { i++ }()
	return 1
}
