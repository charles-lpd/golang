package main

import (
	"fmt"
	"time"
)

func main() {
	// 日期和时间相关函数的方法

	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间是 %v,type=%T \n", now, now)

	// 通过 new 获取年月日，时分秒
	fmt.Printf("年 = %v \n", now.Year())
	fmt.Printf("月 = %v \n", now.Month())      // 4 月 === 'April'
	fmt.Printf("月 = %v \n", int(now.Month())) // 通过 int 转成数字
	fmt.Printf("日 = %v \n", now.Day())
	fmt.Printf("时 = %v \n", now.Hour())
	fmt.Printf("分 = %v \n", now.Minute())
	fmt.Printf("秒 = %v \n", now.Second())

	// 格式化日期时间
	// 第一种
	fmt.Printf("当前年月日是 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 第二种
	dataStr := fmt.Sprintf("当前年月日是 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dataStr = %v", dataStr)

	// 第三种 // 2006-01-02 15:04:05 格式写死不能更改
	// 取全部的时间
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	// 取年月日
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	// 取时分秒
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	// 取月份就是 01，取日就是 02，取时就是 15
	fmt.Printf(now.Format("01 "))
	fmt.Println()

	// 需求，每隔一秒打印一个数字，打印到 100 时推出
	// 需求2 每隔 0。1 秒打印一个数字，打印到 100 时推出
	// 时间的常量
	const (
		Nanosecond  = 1                  // 纳秒
		Microsecond = 1000 * Nanosecond  // 微秒
		Millisecond = 1000 * Microsecond // 毫秒
		Second      = 1000 * Millisecond // 秒
		Minute      = 60 * Second        // 分钟
		Hour        = 60 * Minute        // 小时
	)
	i := 0
	for {
		i++
		fmt.Println(i)
		// time.Sleep(time.Second) // time.Second == 1秒
		time.Sleep(time.Millisecond * 100) // == 100 毫秒
		if i == 5 {
			break
		}
	}

	// Unix 和 UnixNano(纳秒) 的使用
	fmt.Printf("unix时间戳= %v \n, unixnano 纳秒时间戳 = %v \n", now.Unix(), now.UnixNano())

}
