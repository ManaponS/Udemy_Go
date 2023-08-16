package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happen", err)
		//2023/08/16 22:47:38 err happen open no-file.txt: The system
		//cannot find the file specified.
	}
}
