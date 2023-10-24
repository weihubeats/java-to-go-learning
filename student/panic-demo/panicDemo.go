package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("11111.")
	defer func() {
		fmt.Println("22222")
		if p := recover(); p != nil {
			fmt.Printf("4444: %s\n", p)
		}
		fmt.Println("5555")
	}()
	// 引发 panic。
	panic(errors.New("3333"))
	fmt.Println("Exit function main.")
}
