package main

import (
	"fmt"
)

// 演示golang 中指针类型的转换
func main() {
	/*
		1. 值类型，都有对应的指针类型，形式为 *数据类型，比如 int 对应的指针就是 *int,
		float32 对应的指针类型就是 *float32 依此类推
		2. 值类型包括 基本数据类型 int 系列，float 系列，bool，string，数组和结构体 struct
	*/
	// 基本数据类型在内存布局
	var i int = 10
	// i 的内存地址是多少
	// 获取变量的内存地址， 用 & 符号
	fmt.Println("i 的地址=", &i)

	// 指针类型，指针变量存的是一个地址，这个地址指向的空间才是值
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型是 *int
	// 3. ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr = %v \n", ptr)
	fmt.Printf("ptr 的地址= %v \n", &ptr)
	fmt.Printf("ptr 指向的值= %v \n", *ptr)

	// 1. 写一个程序, num变量名，并且获取num的地址打印出来
	var num int = 1000
	fmt.Printf("num 的地址= %v\n", &num)

	// 2. 创造 ptr 获取num指针，并修改num值
	var ptr1 *int = &num
	*ptr1 = 101 //  这里直接是根据指针修改值
	fmt.Printf("num 的值= %v\n", num)
}
