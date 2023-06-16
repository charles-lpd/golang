package main

import "fmt"

// init 函数优先调用
func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("main")
}
