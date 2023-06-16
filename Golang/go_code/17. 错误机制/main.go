package main

import (
	"errors"
	"fmt"
)

// 分母 不能为 0 所以运行会出现错误
//func test() {
//	num1 := 10
//	num2 := 0
//	res := num1 / num2
//	fmt.Printf("res = %v", res)
//}

func test() {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Printf("res = %v", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		// 如果读取文件发送错误。就输出这个错误，并终止程序。
		panic(err)
	}
	fmt.Println("test02() 后执行")
}
func main() {
	// 测试
	//test()
	//
	//fmt.Println("test() 后执行")
	test02()
}
