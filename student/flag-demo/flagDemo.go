package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

// go run flagDemo.go -intflag 12 -boolflag 1
func main() {

	// 解析启动参数
	flag.Parse()
	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)

	var env = flag.String("spring.profiles.active", "default", "this is spring active")
	fmt.Println(*env)

}
