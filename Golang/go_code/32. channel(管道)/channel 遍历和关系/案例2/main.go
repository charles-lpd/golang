package main

import (
	"fmt"
	"time"
)

// 向 intChan 写入 8000个数据
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan<-i
	}
	// 关闭 intChan
	close(intChan)
}

// 从 intChan 取出数据，并判断是否为素数，如果是 就放入到 primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	for {
		num,ok := <- intChan
		if !ok {  // intChan 无法取出值了
			break
		}
		prime := isPrime(num)
		if prime {
			primeChan <- num
			fmt.Println("存入数据=",num)
		}
	}
	fmt.Println("取不到数据 退出。")
	exitChan <- true
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
				return false
		}
	}
	return true
}
func main() {
	
	/*
		需求：要求统计 1-8000 的数字中，哪些是素数？
	*/

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果

	// 标识退出的管道
	exitChan := make(chan bool, 4) // 4 个 协程，就是 4个 bool 来确定完成工作
	start := time.Now().Unix()
	// 开启一个协程， 向 intChan 放入 1-8000 个数
	go putNum(intChan)

	// 开启 四 个协程， 判断是否为素数，放入 primeChan 中
	for i:=1;i<=4;i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 判断 4 个协程完成工作
for {
		if len(exitChan) == 4 {
			end := time.Now().Unix()
			fmt.Println("消耗时间", end - start)
		close(primeChan)
		break
	}
}
// or // 两种都可以
	// 匿名函数 // 判断是否还能取出
	// go func(){
	// 	for i := 0; i < 4; i++ {
	// 		<-exitChan
	// 	}
	// 	close(primeChan)
	// }()


	// 遍历 primeChan
		for {
			// 遍历取出值
			_, ok := <- primeChan

			// 判断 OK 是否取出失败
			if !ok {
				break
			}
			// fmt.Printf("素数 = %v \n",res)
		}
		fmt.Println("main 主线程退出")
}