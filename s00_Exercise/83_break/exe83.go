package main

import "fmt"

func main() {
	i := 0
	for i < 100 {
		if i == 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
