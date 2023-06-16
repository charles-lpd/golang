package main

import (
	"fmt"
)

func main() {
	//使用 for 循环完成一个矩形
	/*
	***
	***
	***
	 */
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 打印 半个金字塔
	/*
	 *
	 **
	 ***
	 */

	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 打印整个金字塔
	/*
		  *
		 ***
		*****

	*/
	totalLevel := 9
	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel+i-1; j++ {
			if j >= totalLevel-i+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
