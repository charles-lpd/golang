package main

import (
	"fmt"
)

type MethodUtils struct {
}
type Calcuator struct {
	Num1, Num2 float64
}

func (z MethodUtils) Rectangular() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (z MethodUtils) Rectangular2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (z MethodUtils) area(len float64, with float64) float64 {
	return len * with
}

func (z MethodUtils) JudgeNum(n int) bool {
	return n%2 == 0
}

func (z MethodUtils) Rectangular3(n, m int, key string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func (calcuator Calcuator) Counting(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}

func main() {
	/*
		1. 编写结构体(MethodUtils),编写一个方法，方法不需要参数
		在方法中打印 一个 10 * 8 的矩形，在main 方法中调用该方法
	*/
	var z MethodUtils
	z.Rectangular()
	/*
	   2. 编写一个方法 提供 m 和 n 两个参数，方法中打印一个 m * n 的矩形

	*/
	fmt.Println()
	z.Rectangular2(10, 8)
	/*
		3. 编写一个方法，计算该矩形的面积(可以接受长len， 和宽 width)，
		将其做为方法返回值，在 main 函数中调用该方法，接受返回的面积值并打印
	*/
	var data = z.area(2.5, 8.7)
	fmt.Println("data = ", data)

	/*
		4. 判断一个数 是奇数还是偶数
	*/
	var data2 = z.JudgeNum(12)
	if data2 {
		fmt.Println("是偶数 = ", data2)
	} else {
		fmt.Println("是奇数 = ", data2)
	}
	/*
		5. 根据行，列，字符打印，对应行数和列数的字符，
			比如 ： 行:3,列:2,字符*,则打印相应的效果
	*/
	z.Rectangular3(2, 3, "$")
	/*
		6. 实现 加减乘除
	*/
	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2
	data3 := calcuator.Counting('+')
	fmt.Println("data3 = ", data3)
}
