package model

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("测试开始")
	t.Run()
}
func TestUser(t *testing.T) {
	fmt.Println("开始测试 user中的相关方法")

	t.Run("测试获取用户数据", testGetUserByUsername)
}

/* func testAddUser(t *testing.T) {
	println("测试添加用户")
	user := &User{}

	// 调用添加用户的方法
	user.AddUser2()
} */

func testGetUserByUsername(t *testing.T) {
	fmt.Println("通过 username 获取数据")
	user := User{
		Username: "charles",
	}
	u, err := user.GetQueryRow()
	if err != nil {
		fmt.Println("获取数据错误")
	}
	fmt.Println(u)
}
