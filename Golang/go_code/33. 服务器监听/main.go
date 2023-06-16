package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	// 这里我们循环的介绍客户端发来的数据
	defer conn.Close() // 一定要关闭

	for {
		// 创造一个新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1. 等待客户端通过 conn 发送信息
		// 2. 如果客户端没有 wrtie 发送，那么协程就阻塞在这里
		fmt.Println("服务器在等待客户端发送信息" + conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从 conn 读取

		if err != nil {
			fmt.Println("服务器的Read err = ", err)
			return // 服务器 读取错误 return 退出
		}
		// 3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
func main() {

	fmt.Println("服务器开始监听。。。。")

	// net.Listen("tcp","0.0.0.0:8888")
	// 1. tcp 表示使用网络协议是 tcp
	// 2. 0.0.0.0:8888 表示 在本地监听 8888 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err = ", err)
	}
	defer listen.Close() // 延时关闭 listen

	// 循环等待客户端来链接我
	for {
		// 等待 客户端链接

		fmt.Println("等待客户端来链接。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
		} else {
			fmt.Printf("Accept() suc con %v,客户端ip = %v \n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}

	// fmt.Printf("listen suc = %v", listen)
}
