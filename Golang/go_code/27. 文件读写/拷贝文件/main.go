package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)
func CopyFile(dstFileName string, srcFileName string)(written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err !=nil {
		fmt.Printf("打开文件失败err = %v", err)
		return 
	}
	defer srcFile.Close() // 关闭文件句柄

	// 通过 srcFile 获取到 Reader
	reader := bufio.NewReader(srcFile)
	// 打开 distFileName

	dstFile, err := os.OpenFile(dstFileName ,os.O_WRONLY | os.O_CREATE,0666)
	if err !=nil {
		fmt.Printf("打开文件失败 err = %v",err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close() // 关闭文件句柄
	return io.Copy(writer,reader)

}
func main(){
	// 拷贝 abc.txt 拷贝到当前文件见下
	srcFile := "../abc.txt"
	dstFile := "./abc.txt"
	_,err := CopyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝失败")
	}

}