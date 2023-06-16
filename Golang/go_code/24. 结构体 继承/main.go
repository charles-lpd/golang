package main

import "fmt"

type Student struct {
	Name  string
	age   int
	Score string
}
type NewStudent struct {
	Student
	AAA int
}

func (stu *Student) setAge(n int) {
	stu.age = n
}
func main() {
	test := NewStudent{}
	// test.age = 123 // 也可以不用写 Student
	test.Student.age = 123
	fmt.Println(test.Student)
	// test.setAge(100)// 也可以不用写 Student
	test.Student.setAge(100)
	fmt.Println(test.Student)

	// 添加默认值
	test2 := NewStudent{
		Student{"", 111, "333"},
		8,
	}
	fmt.Println(test2)
}
