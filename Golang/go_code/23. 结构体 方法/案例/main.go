package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Printf("%v 是一个好人 \n", p.Name)
}

func (p Person) sum() int {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}

	return res
}
func (p Person) sum2(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}

	return res
}
func (p Person) getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	/*
		1. 给 Person 结构体添加 speak 方法，输出 xxx 是一个好人
	*/
	var p Person
	p.Name = "charles"
	p.speak()

	/*
		2. 给 Person 结构体添加计算方法，可以计算从1 + ... +1000的 结果 并返回
	*/

	var data = p.sum()
	fmt.Println("data = ", data)

	/*
		3. 给 Person 结构体添加计算方法，可以从指定位置开始添加到 1 + ... + n 的结果并返回
	*/
	var data2 = p.sum2(100)
	fmt.Println("data2 = ", data2)

	/*
		4. 给 Person 结构体添加 getSum 方法，可以计算两个数的和，并返回结果
	*/

	var data3 = p.getSum(10, 30)
	fmt.Println("data3 = ", data3)
}
