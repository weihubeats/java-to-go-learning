
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



如果要使用结构体的值，还需要将结构体的指针传进来，要实现java上面set方法


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

## 结构体嵌套


## tag

在Go语言中，结构体的tag（标签）是用于给结构体字段附加元数据的一种机制。Tag是一个字符串，可以在结构体字段的类型定义中使用反引号（`）括起来，位于字段的后面。

结构体的tag可以用于多种用途，包括但不限于以下几个方面：

1. 序列化和反序列化：通过在结构体字段上添加`json`或`yaml`标签，可以指定字段在序列化为JSON或YAML格式时的名称和其他选项。

2. 数据校验：通过在结构体字段上添加`validate`标签，可以指定字段的校验规则，例如最大长度、正则表达式等。

3. ORM（对象关系映射）：在使用ORM框架时，可以使用tag来指定数据库表的列名、字段类型、索引等信息。

4. 文档生成：通过在结构体字段上添加`doc`标签，可以为字段生成文档，例如API文档。

```go
type Student struct {
	Name string `yaml:"name"`
	age  int `yaml:"age"`
}
```


这里表示序列化成yaml Name 用name
age 用age

也可以用json
```go
type Teacher struct {
	Name string `json:"name"`
	age int `json:"age"`
}
```

也可以指定多个
```go
type Teacher struct {
	Name string `json:"name" yaml:"name"`
	age int `json:"age" yaml:"age"`
}
```
