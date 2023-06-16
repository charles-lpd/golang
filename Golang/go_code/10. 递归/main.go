package main

import "fmt"

func fbn(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
func main() {
	/*
		斐波那契数列
		请使用递归的方式，求出斐波那契数， 0，1，1，2，3，4，8，13，。。。
		给一个整数 求出它的 斐波那契数是多少
	*/
	fmt.Println(fbn(4))
}
