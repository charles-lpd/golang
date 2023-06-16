package main

import "fmt"

func main() {
	// map 排序 是 key-value 数据结构，又称为字段或者关联数组，
	// 类似其他编程语言的集合，在编程中经常使用到。

	/**
	var map变量名 map 排序[keyType]valueType

	keyType 可以是什么类型
	golang 中的 map的 key 可以是很多种类型，
	比如 bool，数字，string,指针，channel，
	还可以是只包含前面几个类型的接口，22. struct 结构体，数组，
	通常为 int， string
	注意： slice、map 排序 还有 function 不可以， 因为这几个没法用 == 来判断

	valueType 可以是什么类型
	valueType 的类型和key 基本一样
	通常为 数字(整数，浮点数)，string，map 排序 struct
	声明
		var a map 排序[string]string
		var a map 排序[string]int
		var a map 排序[int]string
		var a map 排序[string]map 排序[string]string
	注意：声明是不会分配内存的，初始化需要make，分配哪出年后才能赋值和使用
	*/

	// map 排序 的声明和注意事项
	// 在使用map 前，需要先make，make的作用就是给 map 排序 分配空间
	// map 排序 的 key-value 是无序的

	// 第一种方式
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "刘"
	a["no2"] = "沛"
	// map 排序 的 key 不能重复的，否则会覆盖之前相同的值
	a["no1"] = "东"
	fmt.Println(a)

	// 第二种方式，不指定容量
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "杭州"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 第三种
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "吴用",
		"hero3": "武松",
	}
	//若分配默认 key-value 后以下内容不会报错，在不分配默认值是，是 nil，会报错
	heroes["hero4"] = "打虎"
	fmt.Println(heroes)

	// 保存 3个学生信息，每个学生都有name sex 信息

	studentMap := map[string]map[string]string{
		"liu": {
			"name": "charles",
			"sex":  "男",
		},
		"pei": {
			"name": "charles_dong",
			"sex":  "男",
		},
		"dong": {
			"name": "charles_d",
			"sex":  "男",
		},
	}
	fmt.Println(studentMap)

	/*
		map 是引用类型，遵守引用类型传递的机制，
		在一个函数接受 map，
		修改后会直接修改原来的map
	*/
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 92
	map1[3] = 93
	map1[4] = 94
	map1[5] = 95
	modify(map1)
	// 经过 函数修改 map[1] 被修改为 100 ，所以验证map 是 引用类型
	fmt.Println(map1)
	var map2 []map[int]int
	map2 = make([]map[int]int, 2)
	map2[0] = make(map[int]int, 2)
	map2[0][0] = 0
	map2[0][2] = 1 // make 指定为 2，会自动扩容
	map2[2][0] = 2 // make 指定 map 切片，容量为 2 ，访问越界。报错
	fmt.Println(map2)
	/*
		// 得出结论 map 必需使用 make 来进行扩容，
		// 当 map 是数组时， 索引不能超过 make 指定容量 否则报错
		// 当 map 时 map[type]type 时 make后 可以进行越界，因为会自动扩容
	*/
}

func modify(map1 map[int]int) {
	map1[1] = 100
}
