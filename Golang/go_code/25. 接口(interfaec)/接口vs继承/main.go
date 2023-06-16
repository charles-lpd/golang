package main

import "fmt"

// 创建 Monkey 结构体
type Monkey struct {
	Name string
}
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

// 给 Monkey 添加方法
func (this *Monkey) climBing() {
	fmt.Println(this.Name, "生来会爬树")
}

// 继承 Monkey 结构体
type LittleMonkey struct {
	Monkey
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习会分飞翔了")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习会游泳了")
}

/*
 1. 当 A 结构体继承了B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用
 2. 当 A 结构体需要扩展功能，同时不希望取破坏继承关系，则可以可以取实现某个接口，即可，
    因此我们可以认为，实现接口是对继承机制的补充
*/
func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climBing()
	var b BirdAble = &monkey
	var c FishAble = &monkey
	b.Flying()
	c.Swimming()

}
