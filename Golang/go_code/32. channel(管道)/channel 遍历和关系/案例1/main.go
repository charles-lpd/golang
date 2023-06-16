package main
import (
	"fmt" 
)


// write Data
func writeData(intChan chan int) {
	for i := 0; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData 写入数据 = %v \n", i)
	}
	close(intChan)
}

// read Data
func readData(intChan chan int, exitChan chan bool) {
	for{
		// v 代表 值，ok 代表是否 取值成功
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读取到数据 = %v \n", v)
	}
	// readData 读取数据完成，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {

		/*
		完成 goroutine 和 channel 协同工作的案例
		1. 开启一个 writeData 协程，向管道 intChan 中写入 50 个整数
		2. 开启一个 readData 协程，从管道 intChan 中读取writeData 写入数据
		3. 注意 write 和 readData 操作的是同一个管道
		4. 主线程需要等待writeData 和 readData 协程都工作完成才能退出

		注意 管道 写入数据和取出数据，不分快慢。
		但是写入时必须有 make 的 空间
		读取时， 管道内必需有值
		*/

		// 创建 两个 管道
		intChan := make(chan int, 50) // make 空间 50个 int
		exitChan := make(chan bool, 1) // make 空间 1个 bool

		go writeData(intChan)
		go readData(intChan, exitChan)
		// 以上代码，运行后会直接退出，因为主线程没有任何操作。
		// 所以我们需要让主线程等待 代码运行完成

		for {
			// _, ok := <-exitChan
			// if !ok {
			// 	break
			// }
			if <- exitChan {
				break
			}
		}
		// or
		// <- exitChan
}