package main

import (
	"fmt"
	"strconv"
)

// 演示讲 golang 中的基本数据类型转成string 使用
func main () {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar  byte = 'h'
	var str string

	// golang 中的基本数据类型转成string 使用
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type = %T str= %v \n", str,str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type = %T str= %v \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type = %T str= %v \n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type = %T str= %v \n", str, str)

	// golang 中的string数据类型转成基本数据类型使用
	var str1 string = "true"
	var b1 bool = true
	// b, _ = strconv.ParseBool(str1)
	// 说明
	// 1. strconv.ParseBool(str) 函数回饭回两个值 (value bool, err error)
	// 2. 只想获取 value bool，不想获取err 所以使用 _ 进行忽略
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type %T b1=%v \n", b1, b1)

	var str2 string = "1234560"
	var n1 int64
	var n2 int
	var n3 float64
	var n4 int64
	// ParseInt(s string, base int ,bitSize int)(i int64, err error)
	// s string类型， base 指定进制 2到36 如果 base 为 0 则会从字符串前置判断 "0x"是16进制， "0"是8进制，否则是10进制
	// bitSize 指定结果必须是能无溢出赋值的整数类型， 0，8，16，32，64 分别代表 int, int8, int16, int32, int64
	// 返回的err 是NumErr类型的，如果语法有误，errError = ErrSyntax；如果结果超出范围err.Error = ErrRange
	n1,_ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type = %T, n1=%v \n",n1, n1)
	n2 = int(n1)
	fmt.Printf("n2 type = %T, n2 = %v\n", n2,n2)
	var str3 string = "123.456"
	n3,_ = strconv.ParseFloat(str3, 64)
	fmt.Printf("n3 type = %T, n3 = %v\n", n3, n3)

	// string 转基本类型时， 若不能转成整数，golang 直接转成 0 若是 bool 类型转为 false
	var str4 string = "hello"
	n4, _ = strconv.ParseInt(str4, 10, 64) // n4 =0
	fmt.Printf("n4 tyle = %T, n4 = %v \n", n4, n4)
}