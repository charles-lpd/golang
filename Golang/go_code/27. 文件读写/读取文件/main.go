package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "../test.txt"
	// 判断文件路径是否存在
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File path %v exists.\n", filePath)
	} else if os.IsNotExist(err) {
		fmt.Printf("File path %v does not exist.\n", filePath)
	} else {
		fmt.Printf("Error checking file path %v: %v\n", filePath, err)
	}

	// 使用 os.RedFile 一次性将文件读取
	content,err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件错误，%v",err)
	}
	// 输出读取内容
	// fmt.Printf("%v",content)// 输出的是 []byte，需要转 string 类型
	fmt.Printf("%v",string(content))
	// 文件的 Open 和 Close 被封装到 ReadFile 函数内部
}