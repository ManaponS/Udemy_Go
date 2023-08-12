package main

import (
	"fmt"
	"strconv"
)

type newspaper struct {
	title string
}

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("Title: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("Count: ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Ready Player One",
	}

	var n count = 42

	news := newspaper{
		title: "Multichon",
	}
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println(news)
}
