package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//Panicln is equivalent to Println() followed by a call to panic().

		log.Panicln(err)
	}
}
