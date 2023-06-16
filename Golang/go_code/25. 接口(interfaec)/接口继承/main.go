package main

import "fmt"

type BInterface interface {
	test01()
	test02()
}
type CInterface interface {
	test02()
}

// AInterface 接口继承
type AInterface interface {
	BInterface
	CInterface
	test03()
}
type Student struct {
}

type Stu struct {
}

func (s Student) test01() {
	fmt.Println("test01()")
}
func (s Student) test02() {
	fmt.Println("test02()")
}
func (s Student) test03() {
	fmt.Println("test03()")
}

type A interface {
	say()
}

func (s *Stu) say() {
	fmt.Println("say()")
}
func main() {
	// Student 必需实现 AInterface 中所有定义的方法
	var a AInterface = Student{}
	a.test01()
	a.test02()

	var stu = Stu{}
	// var b A = stu,
	var b A = &stu // 会出现错误，因为 say 接受是的是一个地址 所以需要 &
	b.say()
}
