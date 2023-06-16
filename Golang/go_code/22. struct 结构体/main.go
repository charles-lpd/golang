package main

import "fmt"

// Cat 定义 Cat 结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}
type Person struct {
	Name   string
	Age    string
	Scores [5]float64
	Ptr    *int              // 指针
	slice1 []int             //切片
	map1   map[string]string // map
}

type Monster struct {
	Name string
	List map[int]int
}

func main() {
	/**
	Golang 语言面向对象编程 说明：
	golang 也支持面向对象变成(OOP)，但是和传统的面向对象编程有区别
	并不是纯粹的面向对象语言，所以我们说 Golang 支持 **面向对象编程特性** 是比较准确的

	golang 没有类(class) ，Go 语言的结构体(struct) 和其他编程语言的类(class) 有同等的地位
	你可以理解goLang 是基于 struct 来实现 OOP 特性的。

	Golang 面向对象编程 非常简洁，
	去掉了传统 OOP 语言的集成，方法重载，构造函数和 析构函数，隐藏的 this 指针等等。

	Golang 仍然有面向对象编程的集成，封装 和多态的 特性，只是实现的方式和 其他 OOP 语言不一样
	比如 继承，Golang 没有 extends 关键字，继承是通过匿名字段来实现

	Golang 面向对象(OOP)很优雅，OOP 本身就是语言类型系统(type system) 的一部分，
	通过接口(interface)关联 耦合性低，非常灵活，
	也就是说在 Golang 面向接口编程是非常重要的特性。
	*/

	/*
		老太养了20只猫，一只名字叫小白，今年三岁，白色，还有一只叫小花，
		今年 100岁 ，花色，请编写一个程序，当用户输入小猫名字时，就显示该猫的名字
		年龄，颜色，如果用户输入小猫的名字错误，则显示老太没有这只猫猫

	*/
	// 使用 struct 来完成案例

	// 创建一个 Cat 的变量

	var cat1 = Cat{ // 定义默认属性
		Name:  "小黄",
		Age:   10,
		Color: "黑色",
		Hobby: "玩球",
	}
	//cat1.Name = "小白"
	//cat1.Age = 3
	//cat1.Color = "白色"
	//cat1.Hobby = "吃鱼"

	fmt.Println("cat1 = ", cat1)

	// 如果结构体的字段类型是，指针 slice，和 map的零值都是 nil，即 还没有分配空间
	// 如果需要使用这样的字段，需要先 make 后才能使用
	var map1 = map[string]string{
		"key2": "00000",
	}
	// 需要先分配默认值 key-value,否则 该 map == nil
	map1["key3"] = "12312"
	fmt.Println("map1", map1)

	// 定义 Person 结构体变量
	var p1 Person
	fmt.Println("p1 = ", p1)

	if p1.Ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice1 == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	fmt.Println("map1", map1)

	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.List = make(map[int]int)
	monster1.List[0] = 100
	monster1.List[1] = 200
	fmt.Println("monster1", monster1)
	monster2 := monster1
	monster2.Name = "青牛精"

	// 若 struct 中是引用类型，类似 map 等类型，
	// 需要重新 make 分配空间才可以在进行赋值
	// 否则会改变 monster1 的 值。
	monster2.List = make(map[int]int)
	monster2.List[0] = 300
	monster2.List[1] = 400
	fmt.Println("monster1", monster1)
	fmt.Println("monster2", monster2)

	// 方式1
	var p2 Person
	p2.Name = "tom"
	p2.Age = "18"
	fmt.Println("p2 = ", p2)

	// 方式2
	p3 := Person{Name: "Tom", Age: "100"}
	fmt.Println("p3 = ", p3)

	// 方式3
	var p4 = new(Person)
	(*p4).Name = "123123"
	(*p4).Age = "100"
	fmt.Println("p4 = ", p4)

	// 方式 4
	// var p5 *Person = &Person{Name: "Tom"} // 也可以直接写默认字段
	var p5 = &Person{}
	// 因为 Person 是一个指针，因此标准的访问字段的方法
	// (*person).Name = "tom"
	(*p5).Name = "tom"
	(*p5).Age = "100"
	fmt.Println("p5 = ", p5)
}
