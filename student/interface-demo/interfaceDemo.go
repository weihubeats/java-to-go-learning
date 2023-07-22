package main

import "fmt"

type Fruit interface {

	//方法
	SetName(name string) bool
}

type Apple struct {
	Name string
}

// 实现了 Fruit 接口
func (xiaozou *Apple) SetName(name string) bool {
	xiaozou.Name = name
	return true
}

func main() {
	var apple Apple
	// 给接口赋值
	var fruit Fruit = &apple
	fruit.SetName("苹果")
	fmt.Println("name ", apple.Name)

	// 判断接口是否为 Apple 类型
	if aa, ok := fruit.(*Apple); ok {
		fmt.Println(aa.Name)
	}
}
