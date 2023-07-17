package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main start")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		println("hello A goroutine")
	}()

	go func() {
		defer wg.Done()
		println("hello B goroutine")
	}()
	wg.Wait()
	fmt.Println("main end")

}
