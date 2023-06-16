package main

import (
	"fmt"
)

func main() {
	// 假如还有 97 天放假，问 还剩多少个星期零多少天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d 星期， %d 天 \n", week, day)

	// 定义一个变量保存华氏温度，华氏温度转换摄制温度的公式为：
	// 5/9*(华氏温度 - 100)，请求出华氏温度对应的摄氏温度

	var fahrenheit float32 = 134.2
	var degrees float32 = 5.0 / 9 * (fahrenheit - 100)
	fmt.Printf("%v对应摄氏温度是%v \n", fahrenheit, degrees)

	// for循环练习
	var max int = 100
	var count int = 0
	var sum int = 0
	for i := 1; i < max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v, sum=%v", count, sum)
	fmt.Println("--------------------------")
	const num int = 6
	for i := 0; i < num; i++ {
		fmt.Printf("%v + %v = %v \n", i, num-i, i+num-i)
	}

}
