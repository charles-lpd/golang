package main

import (
	"fmt"
)

func main() {
	/*
		channel 的关闭
		使用内置函数 close 可以关闭 channel， 当 channel 关闭后，
		就不能在向 channel 放数据了，但是仍然可以从 channel 读取数据。

		channel 的遍历

		 channel 支持 for--range 的方式进行 遍历，请注意两个细节
		 1. 在遍历时，如果channel 没有关闭，则会出现 deadlock 的错误
		 2. 在遍历时， 如果channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	*/

	intChan := make(chan int, 3)
	intChan <- 100 // 放入数据
	intChan <- 200 // 放入数据
	close(intChan) // 内置 close 方法， 可以关闭 intChan，关闭后不可放入数据

	// 测试
	// intChan <- 300 // panic: send on closed channel 错误
	fmt.Println("intChan = ", intChan)
	n1 := <-intChan

	fmt.Println("n1 = ", n1)

	// 遍历管道

	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入一百个数据
	}
	fmt.Println("intChan2 = ", intChan2)

	// 遍历 管道 不能使用普通的 for 循环
	// 因为取出时的 intChan2 的 len 发生改变，只能取出五十个
	// for i := 0; i < len(intChan2); i++ {
	// }

	// 在遍历时，如果 channel 没有关闭，遍历完后，则会出现 deadlock 的错误
	close(intChan2) // 需要来关闭管道，遍历后会自动退出， 不会报错
	for v := range intChan2 {
		fmt.Println("v = ", v)
	}

}
