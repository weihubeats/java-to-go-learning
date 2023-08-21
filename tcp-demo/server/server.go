package main

import (
	"fmt"
	"net"
)

const (
	listenAddr = "127.0.0.1:7788"
	network    = "tcp"
)

func main() {

	// 监听指定地址的 TCP 连接
	listener, err := net.Listen(network, listenAddr)
	if err != nil {
		fmt.Println("无法监听地址:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("clone listener error ", err)
		}
	}(listener)

	fmt.Println("等待客户端连接...")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接失败:", err)
			return
		}

		fmt.Println("客户端已连接:", conn.RemoteAddr())

		// 处理连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close Conn error ", err)
		}
	}(conn)

	for {
		// 读取客户端发送的数据
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取客户端数据失败:", err)
			return
		}

		fmt.Println("收到心跳包:", string(buffer))

		// 发送响应给客户端
		_, err = conn.Write([]byte("pong"))
		if err != nil {
			fmt.Println("发送响应给客户端失败:", err)
			return
		}
	}
}
