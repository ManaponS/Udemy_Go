package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("When os.Exit() is called, deferred functions don't run")
	_, err := os.Open("no-file.txt")
	if err != nil {
		//Fatalln is equivalent to Println() followed by a call to os.Exit(1).
		//2023/08/16 22:55:04 open no-file.txt: The system cannot find the file specified. exit status 1
		log.Fatalln(err)
	}
}
