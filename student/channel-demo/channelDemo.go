package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// 开启一个协程
	go Hello(ch)

	// 发送数据到管道中
	ch <- "World"
	//关闭通道
	close(ch)
	time.Sleep(time.Duration(1) * time.Second)

}

func Hello(c chan string) {
	// 从通道获取数据
	name := <-c
	fmt.Println("hello", name)

}
