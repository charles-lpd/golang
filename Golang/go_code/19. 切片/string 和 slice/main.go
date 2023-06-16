package main

import (
	"fmt"
	"strings"
)

func main() {
	// string 的底层是一个 byte 数组，因此 string 也可以进行切片处理

	str := "hello@gmail.com"
	slice := str[6:]
	fmt.Println("slice = ", slice)

	splits := strings.Split(str, string('@'))[1]
	fmt.Println("splits = ", splits)

	// string 是不可变的， str[0] = 'z' 是会错误的
	//str[0] = 'z'

	// 如果需要修改字符串，可以先将 string 转成 []byte 切片 或者 []rune 切片， 修改后，重新转为 string

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str = ", str)

	// 转成 []byte 后 可以处理英文和数字，但是不能处理中文
	// 原因是 []byte 1个 字节来处理，而汉字是3个字节，因此会出现乱码
	// 解决方法 是 将 string 转成 []rune 即可，因为 []rune 是按照字符来处理的，兼容汉字
}
