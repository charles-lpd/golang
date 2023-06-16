package main

import (
	"fmt"
	"sort"
)

// Hero 1. 声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice 2. 声明一个 Hero 结构体的切片类型
type HeroSlice []Hero

// Len 3. 实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 4. Less 方法就是决定你用声明标准的排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}

// Swap 5. 实现 swap 进行交换
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	// 定义一个数组/切片 排序
	slice := []int{0, 2, 5, -1, -100}
	fmt.Println(slice)
	// 使用 sort 包 进行排序
	sort.Ints(slice)
	fmt.Println(slice)

	// 对结构体切片进行排序

	heroes := HeroSlice{
		{
			"charles",
			10,
		},
		{
			"tom",
			18,
		},
		{
			"join",
			31,
		},
	}

	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}
