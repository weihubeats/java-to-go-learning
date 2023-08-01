package main

import "fmt"

func main() {

	var zou = "hello xiao zou"
	fmt.Println(zou)
	ptr := &zou
	// 0x1400010a230
	fmt.Println(ptr)
	// hello xiao zou
	fmt.Println(*ptr)

	*ptr = "update xiao zou"
	// update xiao zou
	fmt.Println(*ptr)
	// update xiao zou
	fmt.Println(zou)

}
