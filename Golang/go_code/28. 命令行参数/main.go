package main

import (
	"fmt"
	"os"
)

// 命令行参数，希望获取到命令行输入的各种参数

func main() {
	// os.Args 是一个 stirng 的切片，用来存储所有命令行参数

	fmt.Println("命令行的参数有", len(os.Args))

	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值

	for i, v := range os.Args {
			fmt.Printf("args[%v]=%v \n",i,v)
	}
}