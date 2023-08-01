## 方法



## 普通方法
```go
func 方法名(参数a, 参数b) (返回值c, 返回值d)  {
	return  "c",true
}

```

```go
func test(a string, b string) (c string, d bool)  {
return  "c",true
}
```

## 类似java的类的方法

注意go中的类是结构体用`struct`关键字指定可以参考
[README.md](student%2Fstruct-demo%2FREADME.md)

```go
func (结构体) 方法名(方法参数) 返回值 {
	// 方法
}
```

```go
func (student *Student) SetName(name string) bool {
	student.name = name
	return true
}
```