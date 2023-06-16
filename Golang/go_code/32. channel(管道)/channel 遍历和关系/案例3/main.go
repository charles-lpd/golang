package main

import (
	"fmt"
)

// 存入数据
func putNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

// 取数据并存入resChan
func getSumResult(numChan chan int, resChan chan int, exitChan chan bool) {
	for {
		res, ok := <-numChan
		if !ok {
			break
		}
		sum := getSum(res)
		resChan <- sum
	}
	fmt.Println("numChan没有数据了")
	exitChan <- true
}
func getSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func main() {
	/*
		1. 启动一个 协程，将 1- 2000 的数放入到 channel 中，比如 numChan
		2. 启动八个协程，从 numChan 取出数 并计算 1+。。。+ n 的值，并存放到 resChan中
		3. 最后 八个协程协同完成工作后，在遍历resChan， 显示结果 如 res[1] = 1 res[10] = 55

	*/

	// 定义Chan  并有 2000 的缓存区
	numChan := make(chan int, 2000)
	// 存放数据的 Chan
	resChan := make(chan int, 2000)
	// 开启 缓存区 为 8 的 ，为确定8个函数时候完成
	exitChan := make(chan bool, 8)
	// 开启协程
	go putNum(numChan)

	// 开启 8个 协程 进行工作
	for i := 1; i <= 8; i++ {
		go getSumResult(numChan, resChan, exitChan)
	}

	for {
		if len(exitChan) == 8 {
			// 如果成功 关闭 resChan 方便遍历，不再写入数据
			close(resChan)
			break
		}
	}
	// 遍历 resChan
	for {
		// 遍历取出值 // 判断是否取出，在8个协程未完成工作时，不能关闭 resChan 否则会导致主线程退出
		res, ok := <-resChan // 若无数据，关闭时走 break ，未关闭时阻塞

		// 判断 OK 是否取出失败
		if !ok {
			break
		}
		fmt.Printf("resChan的值 = %v \n", res)
	}
	fmt.Println("main 主线程退出")
}
