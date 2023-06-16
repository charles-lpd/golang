package main

import "fmt"

// Student 定一个学生结构体
type Student struct {
	Name    string
	Age     int
	Address string
}

func main() {
	/*
		map 的 value 也经常用到 struct 类型，
		更适合管理复杂的数据。
	*/

	// map 的value 经常使用到 struct 类型

	students := make(map[string]Student, 10)

	stu1 := Student{"Tom", 18, "河南"}
	students["no1"] = stu1

	fmt.Println(students["no1"].Name)
	fmt.Println(students)
}
