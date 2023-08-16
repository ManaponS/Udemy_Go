package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//set output for log
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//error will write in log.txt
		log.Println("err happen", err)
	}
	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")
}
