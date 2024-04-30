package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_defer(t *testing.T) {
	defer fmt.Println("小奏技术 start defer 1")
	defer fmt.Println("小奏技术 start defer 2")
	fmt.Println("小奏技术 body")
}

func Test_panic(t *testing.T) {
	panic("an explicit panic")
}

func Test_panic_array_out_of_bounds(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a[3])

}

func Test_panic_gofunc(t *testing.T) {
	// 协程A
	go func() {
		for {
			fmt.Println("goroutine1_print")
		}
	}()

	// 协程B
	go func() {
		time.Sleep(1 * time.Second)
		panic("goroutine2_panic")
	}()

	time.Sleep(4 * time.Second)

}

func Test_panic_defer(t *testing.T) {
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

func Test_panic_for(t *testing.T) {
	x := 42
	fmt.Println(x)

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
