package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	// 判断 users 中用户名是否是唯一
	val, ok := users[name]
	fmt.Println(val, ok)
	if !ok {
		fmt.Println("用户不存在", val, name)
		newUser := make(map[string]string, 2)
		newUser["name"] = name
		newUser["password"] = "123456"
		users[name] = newUser
	} else {
		fmt.Println("用户存在", val, name)
		users[name]["password"] = "88888888"
	}
}
func main() {
	/**
	1. 使用 map[string]map[string]string 的 map 类型
	2. key 表示用户名 是唯一的，不可重复的
	3. 如果 用户存在，就将其密码修改为 "8888" ，如果不存在 就增加这个用户的信息，（包括昵称nickname）
	4. 编写一个函数 modifyUser (users map[string]map[string]string, name string) 完成 上述功能


	*/

	users := make(map[string]map[string]string, 10)
	newUser := make(map[string]string, 2)
	newUser["name"] = "Tom"
	newUser["password"] = "123456"
	users["Tom"] = newUser
	modifyUser(users, "Tom")
	fmt.Println(users)

}
