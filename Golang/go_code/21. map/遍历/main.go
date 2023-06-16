package main

import "fmt"

func main() {
	/**
	相对于复杂的 map 排序 遍历，该 map 排序 的 value 又是一个 map 排序
	说明 map 排序 的遍历使用 for-range 的 结构遍历
	*/

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "杭州"
	cities["no3"] = "上海"
	fmt.Println(cities)
	for key, val := range cities {
		fmt.Printf("key = %v， value = %v \n", key, val)
	}

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

	for key, val := range studentMap {
		fmt.Println("key = ", key, val)
		for key2, val2 := range val {
			fmt.Printf("%v 的 %v = %v \n", key, key2, val2)
		}
	}
}
