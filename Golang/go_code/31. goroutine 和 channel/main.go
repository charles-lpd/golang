package main

/*
	进程
	进程就是程序，程序在操作系统中的一次执行过程，
	是系统进行资源分配和调度的基本单位。

	线程
	线程是进程的一个执行实例，是程序执行的最小单元，
	它是比进程更小的，能独立运行的基本单位。

	一个进程可以创建和销毁多个线程，
	同一个进程中的多个线程，可以进行兵法执行。

	一个程序至少有一个进程，一个进程至少有一个线程
*/

import (
	"fmt"
	"strconv"
	"time"
)

// 编写一个函数，每一秒输出 “hello，world”
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("hellow,word " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	/*
		需求：要求统计1-900000000000 的数字中，哪些是素数？

		分析思路：
		1. 传统的方法，使用循环，判断各数是不是素数
		或
		2. 使用兵法或者并行的方式，将统计素数的任务分配给多个 goroutine 去完成
	*/

	/*

		1. 在主线程中，开启一个 goroutine，该协程每隔1秒输出 “hello，world”
		2. 在主线程中也要每隔一秒输出 "hello,golang",输出10次后，退出程序
		要求主线程和 goroutine 同时执行
	*/
	// test() // 会造后面阻塞，运行完 test() 后才能往后执行
	go test() // 开启了一个协程，会同时执行

	for i := 1; i < 10; i++ {
		fmt.Println("hellow,golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	/*
		hellow,golang 1
		hellow,word 1
		hellow,golang 2
		hellow,word 2
		hellow,word 3
		hellow,golang 3
		hellow,word 4
		hellow,golang 4
		hellow,word 5
		...
	*/
}
