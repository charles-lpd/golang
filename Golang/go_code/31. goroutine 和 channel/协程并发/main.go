package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁 mutex 是互斥
	lock sync.Mutex
)
func test(n int) {

	res := 1
	for i:=1;i<=n;i++ {
		res *= i
	}

	// 加锁
	lock.Lock()
	myMap[n] = res

	// 解锁
	lock.Unlock()
}

func main() {
	// 这里开启多个协程 完成这个任务[200个]
	for i := 1;i<=20;i++ {
		go test(i)
	}

// 休眠十秒钟
time.Sleep(time.Second * 5)

lock.Lock()
	// 输出结果
	for i,v := range myMap {
		fmt.Printf("map[%d] = %d \n",i,v)
	}
lock.Unlock()
}