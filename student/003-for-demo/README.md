## for循环语句

1. 不需要用()括起来
2. 同样支持`continue`和`break`跳出循环


- demo
```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```