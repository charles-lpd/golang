package main

import "fmt"

// AddUpper 闭包就是一个函数和与其相关的应用环境组合的一个整体
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
func main() {
	f := AddUpper()   // 此时  n = 10
	fmt.Println(f(1)) // n = n + 1 = 11 // 并不会被清除
	fmt.Println(f(2)) // 11 + 2 = 13
	fmt.Println(f(3)) // 13 + 3 = 16
}
