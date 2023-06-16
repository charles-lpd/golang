// 表示 该 hello.go 文件所在的包是 main，在 go 中，每个文件都必须归属一个包
package main
//表示引入一个包， 包名称为 fmt 引入改包后可以使用 fmt 的包
import "fmt"

func main()  {
	fmt.Println("hell word!")
}
//使用 go run