
## 简介
个人感觉就是和java的class 类似，把他当java的类理解就行了

## 语法
```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

## 初始化

初始化和groovy有点类似 实际和java 的new 关键词有点类似

```go
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

```

## 结构体的方法

正常java是直接写到类里面
```java
public class Student {
	
	private String name;

	public void setName(String name) {
		this.name = name;
	}
}
```

结构体定义结构体方法是这样的规则
```go
func(结构体名称) 方法名 返回值{
	//具体的方法
}
```



如果要使用结构体的值，还需要将结构体的指针传进来，要实现java类似的方法
```go
type Student struct {
	Name string
	age  int
}

func (student *Student) SetName(name string)  {
	student.Name = name
}
```

## 结构体继承

go 没有继承(extend)关键字 只能组合


