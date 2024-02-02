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

	slice1 := make([]string, 0)

	slice1 = append(slice1, "hahha")
	println("-----")
	//testSlice(&slice1)
	testSlice2(slice1)
	for i := range slice1 {
		println(slice1[i])
	}

	a := [3]int{89, 90, 91}
	modify(a[:])
	for _, value := range a {
		fmt.Println(value)
	}

}

func testSlice(ss *[]string) {
	*ss = append(*ss, "gaga")
}

func testSlice2(ss []string) {
	ss[3] = "gaga"
}

func modify(sls []int) {
	sls[0] = 90
}
