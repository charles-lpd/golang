package main

import (
	_ "fmt"
	"testing"
)

// 编写测试用例，去测试 addUpper 是否正确
// 通过 go test -v 来执行
func TestAddUpper(t *testing.T)  {
	
	res := addUpper(12)
	if res != 55 {
		// fmt.Printf("AddUpper(10) 执行错误，期望值%v, 实际值%v",55,res)
		t.Fatalf("AddUpper(10) 执行错误，期望值%v, 实际值%v",55,res)
	}
	// 如果正确，就 输出日志
	t.Logf("AddUpper(10) 执行正确")
}