package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

func testStruct() {
	 monster := Monster{
		 "牛魔王",
		 500,
		 "2001-09-16",
		 8000.0,
		 "牛魔拳",
	 }

	 // 将 monster 序列化
	 data, err := json.Marshal(&monster)
	 if err != nil {
		 fmt.Printf(" 序列化失败 err = %v",err)
	 }

	 // 输出序列化后的结果 结果是 byte 切片 需要转 string 
	 fmt.Printf(" struct 序列化后的结果是 \n %v",string(data))
}

// 将 map 进行序列化
func testMap() {
	// 定义一个 map

	var a map[string]interface{}

	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"]  = 100
	a["address"] = "不知道"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf(" 序列化失败 err = %v",err)
	}

	// 输出序列化后的结果 结果是 byte 切片 需要转 string 
	fmt.Printf("map 序列化后的结果是 \n %v",string(data))

}

// 将 切片 序列化

func testSlice(){ 
	 var slice []map[string]interface{}
	 var m1 map[string]interface{}
	 // 在使用 map 前需要 make
	 m1 = make(map[string]interface{})
	 m1["name"] = "charles"
	 m1["age"] = 18
	 m1["address"] = "杭州"

	 slice = make([]map[string]interface{}, 2)
	 slice[0] = m1

	 var m2 map[string]interface{}
	 // 在使用 map 前需要 make
	 m2 = make(map[string]interface{})
	 m2["name"] = "charles2"
	 m2["age"] = 16
	 m2["address"] = "河南"

	slice[1] = m2
	// 对 slice 进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf(" 序列化失败 err = %v",err)
	}

	// 输出序列化后的结果 结果是 byte 切片 需要转 string 
	fmt.Printf("slice 序列化后的结果是 \n %v",string(data))
}
func main() {

	// JSON (JavaScript Object Notation) 
	/*
		是一种轻量级的数据交换格式，易于人阅读和编写，同时也易于机器解析和生成， key-val

		JSON 易于机器解析和生成，并有效的提升网络传输效率，
		通常程序在网络传输时会先将数据(结构体，map 等) 序列化成json 字符串
		到接收方得到 json 字符串时，在反序列化恢复成原来的数据类型(结构体，map 等)
		这种方式已然成为各个语言的标准。
	*/


	// 演示将结构体 map 切片，进行序列化
	testStruct()
	testMap()
	testSlice()
} 