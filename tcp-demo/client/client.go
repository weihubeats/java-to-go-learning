package main

import (
	"fmt"
	"net"
	"time"
)

const (
	serverAddr = "127.0.0.1:7788"
	network    = "tcp"
)

func main() {

	// 建立 TCP 连接
	conn, err := net.Dial(network, serverAddr)
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close error ", err)
		}
	}(conn)

	fmt.Println("已连接到服务器")
	// 定时发送心跳包
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		// 获取当前时间
		startTime := time.Now()

		// 发送心跳包数据
		_, err := conn.Write([]byte("heartbeat"))
		if err != nil {
			fmt.Println("发送心跳包失败:", err)
			return
		}

		// 接收服务器响应
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("接收服务器响应失败:", err)
			return
		}

		// 计算时间差
		elapsedTime := time.Since(startTime)

		fmt.Println("收到服务器响应:", string(buffer))
		fmt.Println("心跳花费时间:", elapsedTime)
	}
}
