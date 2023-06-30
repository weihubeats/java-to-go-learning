package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var j = 100.12
	name := "weihubeats"
	name1, name2 := "小奏技术", "ces"
	strNum := strconv.FormatFloat(j, 'f', 2, 64)
	fmt.Print(strconv.Itoa(i) + strNum + name + name1 + name2)
}
