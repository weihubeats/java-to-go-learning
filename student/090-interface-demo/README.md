## 接口定义


type 接口名称 interface {

    method(参数列表) 返回值列表

}

```go

type Student interface {
	
	//方法
	SetName(name string) bool
	
}

```

## 接口的实现
goland的接口实现和java不一样，java需要显现的使用`implements`去实现接口
比如
```java
public class Test implements Ordered
```

goland的接口是隐式的，只要结构体提包含了该方法泽认为该结构体实现接口 例子
```go
// 接口定义
type Fruit interface {

//方法
SetName(name string) bool
}

```