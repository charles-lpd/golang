// 变量的定义和使用
package main
import "fmt"
var (
 c1 int // 默认 0
 c2 = 10000
 c3 bool // 默认false
)
func main()  {
	//  num 变量名， int 变量类型
	// 如果 不赋值 int 的默认值是 0
	var num int
	num = 1000
	fmt.Println("num=", num)

	// num2 根据值自行判断变量类（类型推导）
	var num2 = 10.123
	fmt.Println("num2=", num2)

	// 省略 var， 注意 变量名不可重复声明
	// 下面方式 等价 var name string   name = "liu"
	// := 的 : 不能省略，否则错误
	name := "liu"
	// 重新进行赋值
	name = name + "peidong"

	fmt.Println("name=", name)

	//1. 一次性声明多个变量
	var n1, n2 , n3 int

	fmt.Println("n1=", n1, "n2=", n2, "n3=",n3)

	//2. 一次性声明多个变量 声明 + 赋值
	var a1, a2, a3 = 100, "liu", true

	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)

	// 3 一次性声明多个变量，省略 var + 声明 + 赋值
	b1, b2 , b3 := 100, "peidong", false

	fmt.Println("b1=", b1, "b2=", b2, "b3=", b3)


	// 全局变量 c 

	fmt.Println("c1=", c1, "c2=", c2, "c3=", c3)


}