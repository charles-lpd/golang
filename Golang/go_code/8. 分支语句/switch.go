package main

import "fmt"

func main() {
	var key byte
	fmt.Println("请输入一个字母")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	default:
		fmt.Println("请输入正确字母")
	}
}
