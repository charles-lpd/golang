package main

import "fmt"

func main() {

	// 使用 select 解决从管道取数据的阻塞问题

	// 1. 定义一个管道 10个数据 int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个管道， 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞，导致 deadlock 锁死

	for {
		select {
		// 注意，如果 intChan 一致没有关闭，不会一至阻塞导致 deadlcok 锁死
		case v := <-intChan:
			fmt.Printf("从intchan读取数据 %d \n", v)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取数据 %s \n", v)
		default:
			fmt.Println("读取不到了")
			return
		}
	}
}
