package main

import "fmt"

type Test struct {
	name string
	age  int
}

func main() {
	xiaozou := Test{"sss", 12}
	fmt.Println(xiaozou.name)
}
