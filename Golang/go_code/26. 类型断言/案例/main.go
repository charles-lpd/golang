package main

import "fmt"

// Usb 声明一个接口(interface)
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

// Start 让 Phone 实现 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call() {
	fmt.Println("手机 打电话。。。。")
}

type Camera struct {
	name string
}

// Start 让 Camera 实现 Usb接口的方法
func (p Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (p Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer struct {
}

// 编写一个方法 working，接受 Usb接口类型变量
// 只要是实现了 Usb 接口 (所谓实现了 Usb接口，就是实现了Usb里面声明的方法)
func (c Computer) working(usb Usb) {
	// 通过 usb 接口变量来调用 Start 和 Stop 方法
	// usb 变量会根据传入的实参，判断传入的是 phone 还是 camera // 多态特性
	usb.Start()
	y,ok := usb.(Phone)
	if ok {
		// fmt.Println("执行打电话", y)
		y.Call()
	}
	usb.Stop()
}

func main() {
	// 多态数组
	// 定义一个 Usb 接口数组，可以存放 Phone 和Camear 的结构体变量
	var usbArr [3]Usb
	fmt.Println(usbArr)// [<nil> <nil> <nil>]
	usbArr[0] = Phone{"vivo"} 
	usbArr[1] = Phone{"oppo"}
	usbArr[2] = Camera{"拍立得"}
	fmt.Println(usbArr)// [{vivo} {oppo} {拍立得}]

	/** get Phone 结构体增加一个特有的方法 Call(), 
	当Usb 接口接受的是 Phone 变量时，还需要调用调用 Call 方法
	*/
	// 遍历 usbArr 数组
	var computer Computer

	for _,v :=range usbArr {
		computer.working(v)
		fmt.Println()
	}

}
