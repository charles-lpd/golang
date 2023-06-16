package main

import (
	"fmt"
)

func main() {
	/*
		前面使用 全局变量加锁同步来解决 goroutine 的通讯，但并不完美

		1. 主线程在等待所有 goroutine 全部完成的时间很难确定，我们这里设置 10秒，仅仅是估算

		2. 如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有 goroutine 处于工作状态，这时，也会有随主线程的退出而销毁

		3. 通过 全局变量加锁 同步来实现通讯，也并不利用多个协程对全局变量的读写操作。

		4. 上面的种种分析都在呼唤一个新的通讯机制，channel

	*/

	/*
		1. channel 本质就是一个数据结构 队列

		2. 数据是 先进先出

		3. 线程安全，多 goroutine 访问时 不需要加锁。

		4. channel 是有类型的， 一个 string 的 channel 只能存放 string 类型

	*/

	// 定义 channel 数据

	// var 变量名 chan 数据类型

	// 创建一个可以存放 3个 int 类型的管道

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的 值 %v \n", intChan) //  0xc00011a000 地址

	// 向 intChan 写入数据
	// 在写入数据时 不能超过其容量，也就是 make时 设置的容量
	intChan <- 10
	num := 211
	intChan <- num

	// 看看管道的长度和 cap(容量)
	fmt.Printf("channel len= %v, cap = %v\n", len(intChan), cap(intChan)) // channel len= 2, cap = 3

	// 从 intChan 取数据
	var num2 int

	num2 = <-intChan // 取出第一个数据， 先进先出 原则
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len= %v, cap = %v\n", len(intChan), cap(intChan)) // channel len= 1, cap = 3

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	fmt.Println("num3 = ", num3)

	fmt.Printf("channel len= %v, cap = %v\n", len(intChan), cap(intChan)) // channel len= 0, cap = 3

	// 以下取出会报错
	num4 := <-intChan // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("num4 = ", num4)

	// 注意事项
	/**
	1. channel 中只能存放指定的数据类型
	2. channel 的数据放满后，就不能在放入了
	3. 如果 channel 取出数据后，可以继续存放
	4. 在没有使用协程的情况下，如果我们的 channel 数据已经全部取出，再取就会报告 deadlock
	*/

}
