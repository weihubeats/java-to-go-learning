package main

import (
	"fmt"
	"testing"
)

func TestCreateMap(t *testing.T) {

	// 创建一个map，键类型为string，值类型为int
	var ages map[string]int
	fmt.Println(ages) // 输出：map[]

	// 使用make函数创建map
	ages = make(map[string]int)
	fmt.Println(ages) // 输出：map[]

	// 直接初始化map
	names := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println(names) // 输出：map[Alice:30 Bob:25]

}

func TestUpdateOrDelete(t *testing.T) {

	scores := make(map[string]int)
	scores["Alice"] = 85
	scores["Bob"] = 92

	// 修改map中的元素
	scores["Alice"] = 88

	fmt.Println(scores) // 输出：map[Alice:88 Bob:92]

}

func TestGetValue(t *testing.T) {
	fruits := map[string]float64{
		"Apple":  0.67,
		"Banana": 1.49,
	}

	// 获取元素
	applePrice := fruits["Apple"]
	fmt.Println("Apple Price:", applePrice) // 输出：Apple Price: 0.67

	// 获取不存在的键值
	orangePrice, ok := fruits["Orange"]
	if !ok {
		fmt.Println("Orange is not in the map")
	} else {
		fmt.Println("Orange Price:", orangePrice)
	}

}
