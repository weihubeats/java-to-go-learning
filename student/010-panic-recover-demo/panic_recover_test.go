package main

import (
	"fmt"
	"testing"
)

func Test_defer(t *testing.T) {
	defer fmt.Println("小奏技术 start defer 1")
	defer fmt.Println("小奏技术 start defer 2")
	fmt.Println("小奏技术 body")
}

func TestPanic_defer(t *testing.T) {
	defer fmt.Println("异常退出了")
	fmt.Println("小奏技术 start")
	mockPanic()
	// panic invalid operation: division by zero
	fmt.Println("小奏技术 end")
}

func Test_panic_defer_recover(t *testing.T) {
	recover_panic()
	fmt.Println("小奏技术 end")
}

func recover_panic() {
	defer func() {
		if err := recover(); err != nil {
			// 打印异常，关闭资源，退出此函数
			fmt.Println(err)

		}
	}()
	mockPanic()
}

func mockPanic() {
	a := []int{1, 2, 3}
	fmt.Println(a[3])
}
