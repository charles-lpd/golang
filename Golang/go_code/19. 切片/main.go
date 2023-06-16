package main

import "fmt"

func main() {
	/* 切片的英文是 slice
	 切片是数组的一个应用，因此切片是引用类型，在进行传递时遵守引用传递机制。
	 切片的使用和数组类似，遍历切片, 访问切片的元素和求切片长度 len(slice) 都一样。
	 切片的长度是可以变化的，因此切片是一个可以动态变化的数组
	 切片定义的基本语法:
		var 变量名 []类型
		var name []int
	*/
	intArr := [...]int{1, 22, 33, 44, 55}

	// 声明一个切片
	/*
		1. slice := intArr[1:3]
		2. intArr[1:3] 表示引用到intArr 这个数组
		3. 引用intArr 数组的起始下标为 1， 最后的下标为3(但不包含3)，其实是最后下标 应该是 2
	*/
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice = ", slice)
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("slice 容量= ", cap(slice)) // 切片的容量是可以动态变化的。

	/*
		定义切片的两种方式
	*/
	//  slice := intArr[1:3] 通过引用一个创建好的数组

	// make 方式 make(类型，长度，容量)
	var sliceMake []int = make([]int, 4, 10)
	//
	fmt.Println("sliceMake = ", sliceMake)
	fmt.Println("sliceMake size = ", len(sliceMake))
	fmt.Println("sliceMake cap = ", cap(sliceMake))

	// 第三种 定义一个切片，直接就指定具体数组，使用原理类似于 make 的方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice = ", strSlice)
	fmt.Println("strSlice size = ", len(strSlice))
	fmt.Println("strSlice cap = ", cap(strSlice))

	// 遍历切片
	// 常规 for 循环遍历切片

	// 使用 for-range 方式遍历切片

	/*
		创建切片时的简写方法
		// 从 0 开始取 可以简写省略 0
		var slice = intArr[0:4] => var slice = intArr[:4]

		// 若取截止的数 是数组的最大长度时， 可以省略
		var slice = intArr[1: len(intArr)] => var slice = intArr[1:]

		// 若想直接取全部数组的话， 直接使用 arr[:] 即可
		var slice = intArr[0:len(intArr)] => var slice = intArr[:]

	*/

	// 用 append 内置函数，可以对切片进行动态追加元素
	var slice3 []int = []int{100, 200, 300}

	// 使用 append 直接给 slice3 追加具体元素
	slice3 = append(slice3, 400, 500, 600)
	// append 并不会影响原数组，需重新进行赋值
	fmt.Println("slice3 = ", slice3)

	// 通过 append 将切片 slice3 追加给 slice3
	slice3 = append(slice3, slice3...)

	fmt.Println("slice3 追加切片 slice3 = ", slice3)

	fmt.Println()
	// 切片的拷贝操作
	// 切片使用 copy 内置函数完成拷贝， 举例说明

	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	// 将 slice4 拷贝到 slice5 中
	/*
		copy(para1,para2) 的类型都是 切片类型
		若make 的容量 比 copy 切片小，会出现问题
		copy 一个 切片容量为5 的到 切片容量为 1 的中，
		结果只会拷贝 1 个过去，并不会全部copy过去，因容量为 1。
	*/
	copy(slice5, slice4)

	fmt.Println("slice4 =", slice4)
	fmt.Println("slice5 =", slice5)

	var slice6 []int = []int{3, 4, 5, 6, 7}
	// slice4 覆盖 slice6
	copy(slice6, slice4)
	fmt.Println("slice6 =", slice6)

	// 切片时引用类型，所以在传递时，遵守引用查堵你机制
	test(slice6)
	fmt.Println("slice6 =", slice6)
}

func test(slice []int) {
	slice[0] = 100 // 这里会修改 slice[0], 会改变实参，也就是传递进来的参数
}
