package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.Printf("第%d次循环, x的值为%d \n", i, x)
	}

	strings := [5]string{"aa", "bb", "cc"}
	for i, s := range strings {
		fmt.Printf("第%d次循环, s的值为%s \n", i, s)

	}

	for _, s := range strings {
		fmt.Printf("s的值为%s \n", s)

	}

}
