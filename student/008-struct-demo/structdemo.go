package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `yaml:"name"`
	age  int    `yaml:"age"`
}

func (student *Student) SetName1(name string) {
	student.Name = name
}

// 结构体的方法
func (student *Student) SetName(name string) bool {
	student.Name = name
	return true
}

type Teacher struct {
	Name string `json:"Name" level:"12"`
}

// Response 结构体嵌套结构体
type Response struct {
	Data struct {
		Children []struct {
			Data Student
		}
	}
}

// 继承
type Father struct {
	Name string
}

// 结构体方法
func (father Father) GetName() string {
	return father.Name
}

// 组合Father
type Child struct {
	Father
}

func main() {
	xiaozou := Student{"sss", 12}
	fmt.Println(xiaozou.Name)
	xiaozou.SetName("哈哈")

	xiaozou1 := Student{Name: "小奏", age: 18}
	fmt.Println(xiaozou1.Name)

	// 结构体标签 start
	teacher := Teacher{"老师"}
	//反射获取标签
	typeOf := reflect.TypeOf(teacher)
	name, ok := typeOf.FieldByName("Name")
	if ok {
		fmt.Println(name.Tag.Get("json"), name.Tag.Get("level"))
	}
	// 结构体标签 end

	child := Child{Father{Name: "小奏"}}

	// 两者一样 go这里做了优化 可以直接调用成员变量的方法
	child.Father.GetName()
	getName := child.GetName()
	fmt.Println(getName)

	// 匿名结构体
	anonymizeStruct := struct {
		Name string
	}{
		"匿名结构体",
	}
	fmt.Println(anonymizeStruct.Name)

	// 结构体指针
	var xiaozouPoint *Student

	xiaozouPoint = &Student{Name: "小奏", age: 18}
	fmt.Println(&xiaozouPoint.Name)

}
