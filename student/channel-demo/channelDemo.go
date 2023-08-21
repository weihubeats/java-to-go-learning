package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个 string 类型的chan 容量为1 chan是一个先进先出(FIFO)的队列
	ch := make(chan string, 1)
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

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		// 写入数据
		queue <- i
		fmt.Println("create: ", i)
	}

}
