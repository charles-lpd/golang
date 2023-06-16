package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("输出：", i)
	//}
	//
	//
	// 字符串遍历方式 1- for
	var str string = "hello world!"
	for j := 0; j < len(str); j++ {
		fmt.Printf("%c \n", str[j])
	}
	var str2 string = "hello world!北京"
	// 字符串遍历方式 2 for-range
	for index, val := range str2 {
		// 值和下标
		fmt.Printf("%c %v\n", val, index)
	}
}
