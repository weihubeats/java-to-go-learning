package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	configString := string(file)
	fmt.Println("配置文件内容:", configString)
}
