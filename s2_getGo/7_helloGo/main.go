package main

import "fmt"

func main() {
	var name, age = "Sukmai", 45
	emoji := "🐱‍👤"
	fmt.Printf("My name is %v \nI'm now %v \nboth is %T and %T", name, age, name, age)
	fmt.Printf("%v type: %T", emoji, emoji)
}
