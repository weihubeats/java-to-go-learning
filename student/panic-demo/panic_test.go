package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	// 使用 defer 来注册异常处理函数
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常:", err)
		}
	}()

	// 执行可能出现异常的代码
	DoSomething()

	// DoSomething() 后面的代码会得到执行
	fmt.Println("后续代码")

}

func DoSomething() {
	// 使用单独的 defer 语句来触发异常
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	// 模拟出现异常
	panic("出现异常")
}
