package main

import "fmt"

// 在函数中，程序员经常需要创建资源（比如，数据库、文件句柄，锁等），
// 为了在函数执行完毕后，及时释放资源，Go的设计者提供 defer(延时机制)

func sum(n1, n2 int) int {
	// 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈中
	// 当函数执行完毕后，再从 栈中 ，按照先入后出的方式出栈，执行
	defer fmt.Println("n1=", n1) // 1.先入后出
	defer fmt.Println("n2=", n2) // 2.先入后出
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res = ", res)
	return res

}
func main() {
	res := sum(10, 20)
	fmt.Println("return res=", res)
	var i = 0
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
}
