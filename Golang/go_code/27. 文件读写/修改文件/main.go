package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	//打开文件，并追加语句

	// 1. 打开文件
	filePath := "../abc.txt"
	file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND ,0666)
	if err != nil {
		fmt.Print("打开文件错误", err)
		return
	}
	// 及时关闭 file 句柄
	defer file.Close()

	// 准备写入 5 句 “hello, charles”
	str := "我不好, charles\n"

	// 写入时，使用带缓存的 *Writer
	writer :=bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	
	/**
		因为writer 是带缓存，因此在调用 WriteString 方法时
		内容是先写入到缓存的，所以需要调用 Flush 方法，
		将缓存的数据真正的写入到文件中，否则文件中会没有数据
	*/
	writer.Flush()

}