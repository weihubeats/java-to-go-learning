package main

type Student struct {
	name string
}

func main() {

}

func test(a string, b string) (c string, d bool) {
	return "c", true
}

func (student *Student) SetName(name string) bool {
	student.name = name
	return true
}
