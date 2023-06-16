package main

import "fmt"
import "unsafe"
func main() {
	// 布尔值
	var b bool = true
	fmt.Println("bool=", b)

	// 数字类型
	// 整数 int 浮点型 float32， float64, Go 语言支持类型和浮点型数字
	// 并且支持复数，其中位的运算采用补码
	/*
		 Go 也有基于架构的类型 例如 int uint 和 uintptr
		 uint8 无符号 8位整数（0 到 255）
		 uint16 无符号 16位整数（0 到 65535）
		 uint32 无符号 32 位整型 (0 到 4294967295)
		 uint64 无符号 64 位整型 (0 到 18446744073709551615)

		 int8 有符号 8 位整型 (-128 到 127)
		 int16 有符号 16 位整型 (-32768 到 32767)
		 int32 有符号 32 位整型 (-2147483648 到 2147483647)
		 int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

		 float32 IEEE-754 32位浮点型数
		 float64 IEEE-754 64位浮点型数
		 complex64  32 位实数和虚数
		 complex128 64 位实数和虚数

		 byte 类似 uint8
		 rune 类似 int32
		 uint 32 或者64 位
		 int 与 uint 一样大小
		 uintptr 无符号整型，用于存放一个指针

	*/
	var i int = 100
	fmt.Println("int=", i)
	var i32 float32 = 100.000001
	fmt.Println("int32=", i32)

	// 字符串类型
	var s string = "liu"
	fmt.Println("string=", s)

	// 派生类型
	/*
		包括
		指针类型（Pointer）
		数组类型
		结构化类型（struct）
		Channel类型
		函数类型
		切片类型
		接口类型 （interface）
		Map 类型
	*/
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占用的字节数是 %d ",n2, unsafe.Sizeof(n2))
}