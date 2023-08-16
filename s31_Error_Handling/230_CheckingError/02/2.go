package main

import "fmt"

func main() {
	var name int
	fmt.Print("Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}
