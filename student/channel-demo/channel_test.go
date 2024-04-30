package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_channel_no_buffer(t *testing.T) {
	noCache := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("写入数据 ", i)
			noCache <- fmt.Sprintln("小奏写入数据 ", i)
		}
	}()
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		v := <-noCache
		fmt.Println("接受到的数据为: ", v)
	}

}

/*
*有缓冲 channel 的内部有一个缓冲队列
发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间
接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素
*/
func Test_channel_buffer(t *testing.T) {
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))
}

/*
*
channel 使用 close进行关闭，关闭后的channel可以获取数据，但不能继续写入数据
*/
func Test_channel_close(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 2
	ch <- 3
	close(ch)
	v := <-ch
	fmt.Println("get channel data ", v)
	// painc
	ch <- 2

}

func Test_channel_one_dimensional(t *testing.T) {

	ch := make(chan int)
	go producer(ch) // 启动生产者goroutine
	consumer(ch)    // 在主goroutine中运行消费者

}

// 生产者函数，只能向channel中发送数据
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("生产者发送数据: %d\n", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

// 消费者函数，只能从channel接收数据
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("消费者接收数据: %d\n", value)
	}
}

func Test_channel_select(t *testing.T) {

	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}

}

func downloadFile(s string) string {
	intn := rand.Intn(5)
	time.Sleep(time.Duration(intn) * time.Second)
	return s

}
