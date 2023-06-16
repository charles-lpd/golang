package main

import (
	"fmt"
)

// 冒泡排序
func arrSort(arr *[5]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println("arr = ", arr)
}

// 二分查找
func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		// 说明我们要查找的数， 应该在 leftIndex --- middle
		binaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 说明我们要查找的数 应该在 middle + 1 --- rightIndex
		binaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为%v", middle)
	}
}

func main() {
	// 定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	// 冒泡排序
	arrSort(&arr)

	// 在 golang中 我们常用的查找方式有两种
	// 1. 顺序查找
	// 2. 二分查早

	// 从键盘输入一个名称，判断数列中是否包含此名称 [顺序查找]
	names := [4]string{"白眉鹰王", "金毛狮王", "紫杉龙王", "青翼蝠王"}
	heroName := ""
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&heroName)

	// 顺序查找 第一种
	//for i := 0; i < len(names); i++ {
	//	if heroName == names[i] {
	//		fmt.Printf("找到%v 下标是 %v", heroName, i)
	//	} else if i == (len(names) - 1) {
	//		fmt.Printf("没有找到%v", heroName)
	//	}
	//
	//}

	//顺序查找 第二种
	index := -1

	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v 下标是 %v", heroName, index)
	}

	// 二分查找
	arr2 := [6]int{1, 8, 10, 89, 1000, 1234}
	binaryFind(&arr2, 0, len(arr2)-1, -1)
}
