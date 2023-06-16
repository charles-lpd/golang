package main

import (
	"fmt"
	"flag"
)

func main() {

	// 使用 flag 包来解析命令行参数，
	/*
		比 os.Args 要方便很多，而且顺序可以随意
		例如
		-u charles -p 123456
		// -u == 账户
		// -p == 密码
	*/

	// 定义几个变量， 用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	/*
		&user 就是接收用户命令行中输入的 -u 后面的参数值
		"u" 就是 -u 后的指定参数
		"" 默认值是 空字符串
		"用户名默认为空" 说明
	*/
	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&pwd,"p","","密码默认为空")
	flag.StringVar(&host,"h","localhost","主机名默认localhost")
	flag.IntVar(&port,"port",3306, "端口号,默认为 3306")

	// 这里非常重要，转换，必需调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v",user,pwd,host,port)
}
