package main

import "fmt"

type Student struct {
	name string
}

func main() {
	// 匿名函数
	greet := func(name string) {
		fmt.Println("Hello, " + name + "!")
	}
	greet("小奏技术")

}

func test(a string, b string) (c string, d bool) {
	return "c", true
}

func (student *Student) SetName(name string) bool {
	student.name = name
	return true
}
