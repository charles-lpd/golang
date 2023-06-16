package main

import (
	"./model"
	"fmt"
)

func main() {
	fmt.Println(model.HerName)
	model.MainName()
	// 值类型和应用类型
	// 值类型: 基本类型 int 系列，float 系列 bool string 数组和结构体struct
	// 变量直接储存值， 内存通常在栈中分配

	// 引用类型: 指针，slice 切片，map 排序，管道chan，interface 等都是引用类型
	// 引用类型: 变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），
	// 内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据就成为了一个垃圾，由 GC 来回收
}
