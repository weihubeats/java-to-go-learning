# java-to-go-learning
记录go学习中一些语法和demo，也包含一些与java的对比

## github 官方仓库地址
- [go github](https://github.com/golang/go)

## 语法

### 编译
```shell
go.build
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


