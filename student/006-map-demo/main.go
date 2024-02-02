package main

import "fmt"

func main() {
	m := make(map[string]int, 10)
	ss := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	m["weihubeats"] = 1
	m["hahah"] = 2
	m["vv"] = 3
	for site := range m {
		fmt.Println(site, "sss", m[site])
	}

	for k, v := range ss {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除key
	delete(m, "vv")
	// 获取map长度
	len := len(m)
	fmt.Println(len)

	mm := make(map[string]string)
	mm["kkk"] = "value"
	for key := range mm {
		fmt.Println("key ", key, mm[key])
	}

	s, ok := mm["kk"]
	if ok {
		fmt.Println(s)
	}

}
