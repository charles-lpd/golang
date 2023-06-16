package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	// 稀疏数组

	// 1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	// 2. 输出原始数组

	for _, v := range chessMap {
		for _, j := range v {
			fmt.Printf("%d \t", j)
		}
		fmt.Println()
	}

	// 3。 转成稀疏数组
	// 思路
	// 1. 遍历 chessMap 如果我们发现有一个元素的值不为 0， 创建一个 Node 结构体
	// 2. 将其放入到对应的切片中即可

	var sparseArr []ValNode

	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 若不等于 0 ，创建一个 ValNode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	fmt.Println("输出稀疏数组")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d \n", i, valNode.row, valNode.col, valNode.val)
	}

	// 将稀疏数组转变成 原始数组
	var chessMap2 [11][11]int

	// 遍历 sparseArr
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	// 查看 chessMap2 
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d \t", v2)
		}
		fmt.Println()
	}

}
