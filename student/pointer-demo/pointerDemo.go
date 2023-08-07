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

	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

}

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}
