package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("您的年龄大于18，已具备法律责任，要对自己的行为负责！")
	}
}
