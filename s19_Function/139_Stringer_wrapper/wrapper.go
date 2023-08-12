package main

import (
	"fmt"
	"log"
	"strconv"
)

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

func logInfo(s fmt.Stringer) {
	log.Println("Log from 138", s.String())
}

func main() {
	b := book{
		title: "Ready Player One",
	}

	var n count = 42

	logInfo(b)
	logInfo(n)
}
