package main

import "fmt"

func main() {

	var xiaozou []int
	// 添加元数
	xiaozou = append(xiaozou, 1)
	xiaozou = append(xiaozou, 2)
	//遍历
	for index, value := range xiaozou {
		fmt.Print(index, value)
	}

}
