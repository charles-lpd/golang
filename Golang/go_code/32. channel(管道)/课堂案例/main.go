package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}
func main() {
	// 定义一个存放任意数据类型的管道
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom"
	cat := Cat{"小花猫",4}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat = %T， newCat = %v \n",newCat,newCat)
	// fmt.Printf("newCat.Name = %v",newCat.Name) // 编译器不通过， 因为 chan 存放的是 interface{} 类型 所以需要类型断言后使用

	fmt.Printf("newCat.Name = %v",(newCat.(Cat)).Name) // 使用类型断言

}