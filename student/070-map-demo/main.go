package main

import "fmt"

func main() {
	m := make(map[string]int, 10)
	m["weihubeats"] = 1
	m["hahah"] = 2
	m["vv"] = 3
	for site := range m {
		fmt.Println(site, "sss", m[site])
	}

	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除key
	delete(m, "vv")
	// 获取map长度
	len := len(m)
	fmt.Println(len)
}
