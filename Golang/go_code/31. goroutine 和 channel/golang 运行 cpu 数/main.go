package main

import (
	"fmt"
	"runtime"
)
func main() {

	// go v1.18 版本及以上不用设置
	// 以下版本需要设置 cpu 数量达到高效

	cpuNum := runtime.NumCPU()
	fmt.Println("cpunum = ", cpuNum)

	// 可以设置自己使用多少个 CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("cpuNum =", cpuNum)
}