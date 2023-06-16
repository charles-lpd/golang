package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}
func main() {
	// 类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
	var a interface{}
	point := Point{1,2}
	a = point // ok
	var b Point
	// b = a 不可以
	// b = a.(Point) // 可以
	b = a.(Point) 
	// 类型断言， 表示判断 a 是否只想 Point 类型的变量，
	// 如果是就转成 Point 类型并赋值 b 变量，否则报错。
	fmt.Println(b)
	fmt.Println(a)

	// 类型断言，(带检测)
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口，可以接受任意类型
	y,ok := x.(float64) // 类型断言
	if ok {
		fmt.Println(y)
	} else {
		fmt.Println("类型断言失败")
	}
	fmt.Println("继续执行以下代码")

}