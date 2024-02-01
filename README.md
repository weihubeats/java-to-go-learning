# java-to-go-learning
记录go学习中一些语法和demo，也包含一些与java的对比

## github 官方仓库地址
- [go github](https://github.com/golang/go)

## 语法

### 编译
```shell
go build test.go
```
### 运行
编译完后出现一个xxx.exe
直接运行即可
也可以输入
```shell
go run hello.go
```

### 编译规则
1. Go编译器是一行一行进行编译的，多条语句写在一行直接报错
2. Go语言定义的变量或者import导入的包没有使用也会编译报错

### hello word

- java
```java
public static void main(String[] args) {
        System.out.println("hello word");
    }
```

- go
```go
func main() {
	fmt.Println("hello word")
}
```


### 变量定义

- java
```java
String a = "test";
```

- go 
1. 声明变量不赋值
```go
// 使用默认值 0
var i int
```
2. 类型推断不声明类型
```go
var j = 100.12
name := "weihubeats"
```
3.多个变量一起声明
```go
name1, name2 := "小奏技术", "ces"
```
> 相关源码在 student/variantdemo/variantdemo.go


## 值类型和引用类型
值类型: 基本数据类型int、float、bool、string、数组和结构体(struct)
引用类型: 指针、切片(slice)、map、管道(chan)、接口(interface)

## map
[README.md](student%2Fmap-demo%2FREADME.md)

## defer
[README.md](student%2Fdefer-demo%2FREADME.md)

## 结构体
[README.md](student%2Fstruct-demo%2FREADME.md)


## 依赖管理
现在最新最流行的依赖管理是go mod 和java 的maven 的pom类似
文件名: go.mod
使用方式:
github地址 版本

例子: 
`github.com/go-redis/redis v0.0.0-20190803144825-742f3ccb21cd`

`go get github.com/apolloconfig/agollo/v4@v4.3.1`



他的版本不同于maven 需要发到中央仓库，是直接使用git 的tag作为版本号
也可以直接使用指定分支的版本
使用方式
1. 执行命令: `go get  git地址@分支名`
2. 直接在go.mod文件中添加，格式如下：`git地址  v0.0.0-时间戳- commint id`

如果404可以设置代理
```go
go env -w GOPROXY=https://goproxy.cn,direct
```


## 整理一下golang学习资料

- [100天精通Golang(基础入门篇](https://blog.csdn.net/qq_44866828/category_12339137.html)
- [扫清go语言一切障碍，go语言实战](https://github.com/golang-minibear2333/golang)
- [Uber Go 语言编码规范中文版](https://github.com/xxjwxc/uber_go_guide_cn)
- [Goland书籍列表](https://github.com/dariubs/GoBooks)
- [Go语言标准库](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/)
- [go语言圣经](http://www.gopl.io/)
- [go语言圣经中文版](https://golang-china.github.io/gopl-zh/index.html)