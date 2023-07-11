package redisdemo

import (
	"fmt"
	"time"
)

func main() {

	get := Rdb.Get("testKey")
	val := get.Val()
	fmt.Println("redis string value ", val)

	Rdb.Set("testKey", "testValue", 30*time.Second)
}
