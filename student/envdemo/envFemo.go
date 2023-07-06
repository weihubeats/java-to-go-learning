package main

import (
	"flag"
	"fmt"
)

func main() {
	// 启动脚本  go run envFemo.go -spring.profiles.active="prd"
	var env = flag.String("spring.profiles.active", "default", "this is spring active")
	// 解析启动参数
	flag.Parse()
	fmt.Println(*env)

}
