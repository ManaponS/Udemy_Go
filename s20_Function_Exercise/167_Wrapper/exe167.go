package main

import (
	"fmt"
	"time"
)

/*a wrapper function wraps or modifies another function's behavior, while a
callback function is a function passed as an argument to be executed at a specific point or
event.*/

func main() {
	timeFunc(doWork)

	wrappedF := func() {
		fmt.Println("Hello,World")
	}
	logger(wrappedF)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func logger(f func()) {
	fmt.Println("Start")
	f()
	fmt.Println("End")
}
