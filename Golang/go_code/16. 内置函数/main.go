package main

import "fmt"

func main() {
	// len 用来求长度，比如 string， array，mao channel

	// new  // 用来分配内存，主要用来分配值类型，比如 int, float32, struct 返回的时指针
	num1 := 100
	fmt.Printf("num1 的类型是%T, num1 的值是 %v， num1 的地址是%v \n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2 的类型是%T, num2 的值是 %v， num2 的地址是%v", num2, num2, &num2)

	// make // 用来分配内存，主要用来分配引用类型，比如 channel， map 排序， slice。
}
