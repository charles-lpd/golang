package main

import "fmt"

func main() {
	/*

		切片的数据类型如果是 map 排序， 则我们称之为 slice of map 排序 ,map 排序 切片，
		这样使用则 map 排序 个数就可以动态变化了

		案例需求
		使用一个 map 排序 来记录 monster 的信息 name 和age 也就是说一个 monster 对应一个 map 排序
		并且妖怪的个数可以动态的增加 => map切片。
	*/
	// 1.声明一个 map 排序 切片
	//var monster = map 排序[string]string 预估 放 2 个
	monster := make([]map[string]string, 2)
	// 2 增加 第一个妖怪的信息
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "500"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "红孩儿"
		monster[1]["age"] = "100"
	}
	// 超过 两个 以下代码运行会越界。
	//if monster[2] == nil {
	//	monster[2] = make(map[string]string, 2)
	//	monster[2]["name"] = "狐狸精"
	//	monster[2]["age"] = "1000"
	//}

	// 我们需要使用到 切片的 append 函数，可以动态的增加 monster
	// 1。 先定义 monster 信息
	newMonster := map[string]string{
		"name": "新妖怪",
		"age":  "300",
	}
	// 通过 append 插入即可
	monster = append(monster, newMonster)
	fmt.Println(monster)
}
