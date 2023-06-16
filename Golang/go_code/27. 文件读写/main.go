package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	// 通过 os 包 Open 函数 打开文件
	file,err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("打开文件失败",err)
	}
	fmt.Printf("file = %v \n",file)

	// 创建一个 *Reader ，是带缓冲的
	/*
		const(
			defaultBufSize = 4096 //  默认的缓冲区为 4096
		)
	*/
	reader := bufio.NewReader(file)

	// 循环的读取文件的内容
		for{
			str,err := reader.ReadString('\n')
			if err != nil && err !=io.EOF {
				panic(err)
			}
			//输出内容
			fmt.Print(str)
			if err == io.EOF { //io.EOF 表示文件的末尾
				break
			}
		}
	fmt.Println("读取文件结束")
	// 当函数退出时，要及时关闭file，关闭文件
	err = file.Close() // 否则会有内存泄漏
	if err != nil {
		fmt.Println("关闭文件失败",err)
	}

}