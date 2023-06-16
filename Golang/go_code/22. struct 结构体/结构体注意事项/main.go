package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp    Point
	rightDown Point
	// 相同类型可用简写为下面方式
	// leftUp, rightDown Point
}

type A struct {
	Num int
}
type B struct {
	Num int
}

// Monster struct 的每个字段上，可以写上一个 tag ，该 tag 可以通过反射机制获取，常见的使用场景就是序列化和反序列化
// Monster 必需要大写，否则 json.Marshal 无法使用内部变量
type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag，这样在 json化后会变成小写。
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 有四个 int, 在内存中是连续分布的
	fmt.Println(r1)

	// 结构体是用户单独定义的类型，和其他类型进行转换时，
	// 需要有完全相同的字段，个数和类型
	var a = A{100}
	var b = B{300}
	// b = a // 错误
	b = B(a) // 可以使用强行转换进行，
	fmt.Println(b)

	// 创建一个 Monster 变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	// json.Marshal 函数中使用到了反射
	// 返回两个变量， 值 & error
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 序列化失败", err)
	}
	// 打了 tag 前 {"Name":"牛魔王","Age":500,"Skill":"芭蕉扇"}
	fmt.Println("jsonMonster = ", string(jsonMonster))
	// 打了 tag 后 {"name":"牛魔王","age":500,"skill":"芭蕉扇"}
	fmt.Println("jsonMonster = ", string(jsonMonster))

}
