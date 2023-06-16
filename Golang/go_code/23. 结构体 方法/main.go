package main

import "fmt"

type Person struct {
	Name string
}

// 给 Person 绑定 test 方法，只适用于 Person 类型的变量来调用， 其他不可直接调用
func (p Person) test() {
	// 接受 Person 类型的变量
	p.Name = "jack"
	fmt.Println("test() name = ", p.Name)

}
func main() {
	var p Person
	// 调用 test 方法，触发 绑定给 Person 类型的方法，p 调用相当于将 p 传递给 test 方法
	p.test()

}
