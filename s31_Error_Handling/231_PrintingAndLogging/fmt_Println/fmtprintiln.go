package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happen", err)
		//err happen open no-file.txt: The system cannot find the file specified.

	}
}
