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

// 演示将 json 字符串，反序列化成 struct 或 map 或 切片
func main(){
	// json 数据
	str := "{\"name\": \"牛魔王\",\"age\": 500,\"birthday\": \"2001-09-16\",\"sal\": 8000,\"skill\": \"牛魔拳\"}"

	// 定义一个 Monster 示例
	var monster Monster

	err := json.Unmarshal([]byte(str),&monster)

if err != nil {
	fmt.Println("反序列化失败", err)
}
fmt.Println(monster,monster.Name)
}