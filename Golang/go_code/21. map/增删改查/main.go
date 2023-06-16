package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = ""
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 因为 no3 的 key 已经存在，不存在就是增加，存在就是修改
	cities["no3"] = "河南"
	fmt.Println(cities)

	/*
			删除 func delete(m map 排序[Type]Type1, key Type)
			若存在则删除， 若不存在也不会报错 delete(map 排序, "key")

		细节说明
			· 如果我们要删除 map 排序 里所有的 key，没有一个专门的方法去一次删除
			  可以遍历一下 key，逐个删除
			· 或者 map 排序 = make(...)，将内容指向一个新的 make，让原来的成为垃圾，被 gc 回收
	*/

	delete(cities, "no1")
	fmt.Println(cities)

	// 查找
	val, ok := cities["no1"]
	if ok {
		fmt.Println("有 no2 的 key", val, ok)
	} else {
		fmt.Println("没有 no2 的 key", val, ok)
	}
}
