package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin GoRoutine: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("A ", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("B ", i)
		}
		wg.Done()
	}()
	fmt.Println("Mid GoRoutine: ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End GoRoutine: ", runtime.NumGoroutine())
}
