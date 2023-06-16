package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 定义数组
	num1 := [...]int{1, 2, 3, 5, 6}
	fmt.Println("num1 = ", num1)

	num2 := [5]int{1, 2, 3, 5, 6}
	fmt.Println("num2 = ", num2)

	num3 := [...]int{1: 100, 2: 200, 3: 300}
	fmt.Println("num1 = ", num3)

	strArr := [...]string{"1", "2", "3"}
	fmt.Println("num1 = ", strArr)

	// for-range 结构遍历数组
	heroes := [...]string{"刘", "沛", "东"}
	for index, val := range heroes {
		fmt.Printf("index = %v, value = %v \n", index, val)
	}

	// 创建一个 byte类型 的26个元素的数组，分别放置 'A - Z',使用for循环访问所有元素并打印出来， 提示字符数据运算 'A' + 1 => 'B'
	fmt.Printf("值 = %v \n", string('A'+0))

	letter := [26]byte{}
	for index := range letter {
		letter[index] = 'A' + byte(index)
	}
	for _, val := range letter {
		fmt.Printf("%c ", val)
	}
	fmt.Println()
	// 请求出 一个数的最大值，并得到对应的下标
	intArr := [...]int{1, -7, 3, 4, 100, -1}
	count := 0
	maxIndex := 0

	for index, val := range intArr {
		if val > count {
			count = val
			maxIndex = index
		}
	}
	fmt.Printf("最大值是%v, 其下标是%v \n", count, maxIndex)

	// 请求出一个数组的和 和平均值 for-range

	intArr2 := [...]int{1, -1, 9, 90, 12}
	sum := 0
	averageValue := 0.00

	for _, val := range intArr2 {
		sum = sum + val
	}
	// 将 sum 转成 float 64 ,因为会涉及到小数。
	averageValue = float64(sum) / float64(len(intArr2))
	fmt.Printf("数组和是%v, 平均值是%v \n", sum, averageValue)

	// 要求生成五个随机数，并将其反转并打印
	intArr3 := [5]int{}
	for index := range intArr3 {
		num, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			panic(err)
		}
		intArr3[index] = int(num.Int64())
	}
	//var a = [6]int{81, 87, 47, 59, 81, 67}
	fmt.Println(intArr3)
	for i := 0; i < len(intArr3)/2; i++ {
		temp := intArr3[i]
		intArr3[i] = intArr3[len(intArr3)-i-1]
		intArr3[len(intArr3)-i-1] = temp
	}
	fmt.Println(intArr3)

	// 排序 大 -> 小
	var intArr4 = [...]int{15, 17, 90, 4, 1, 0}
	for i := 0; i < len(intArr4)-1; i++ {
		for j := 0; j < len(intArr4)-i-1; j++ {
			if intArr4[j] < intArr4[j+1] {
				intArr4[j], intArr4[j+1] = intArr4[j+1], intArr4[j]
			}
		}
	}
	fmt.Println(intArr4)
}
