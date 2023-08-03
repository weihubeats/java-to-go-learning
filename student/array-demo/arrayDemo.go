package main

import "fmt"

func main() {

	s := [4]string{"aa", "bb"}
	s[3] = "cc"
	fmt.Println(s)

	// 声明数组存的指针
	s1 := [3]*string{new(string), new(string), new(string)}
	*s1[0] = "aa"
	*s1[1] = "bb"

	for _, s := range s1 {
		fmt.Println(*s)
	}

	// 多维数组
	s2 := [4][2]int{{10, 11}, {20, 21}}
	for _, row := range s2 {
		for _, value := range row {
			fmt.Println(value)
		}
	}

}
