package main

import "fmt"

func main() {
	foo()
	bar("BAR")
	s := ret("RETURN")
	fmt.Println(s)
	msg, age := dogYears("Todd", 51)
	fmt.Println(msg, age)
}

// func (receiver) identifier(parameters) (returns) { code }
func foo() {
	fmt.Println("FOO")
}

func bar(s string) {
	fmt.Println("BAR receive", s)
}

func ret(s string) string {
	return fmt.Sprint("RET Return ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprintf("%v is %v year old in dog year", name, age), age
}
