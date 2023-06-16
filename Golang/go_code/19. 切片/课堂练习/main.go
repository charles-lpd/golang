package main

import "fmt"

func main() {
	/*
			1. 可以接受一个 n int
			2. 能够将 斐波那契的数列放到切片中
			3. 提示 斐波那契数列形式
		arr[0] = 1;arr[1] = 1;arr[2] = 2;arr[3] = 3;arr[4] = 5;arr[5] = 8;

	*/
	data := fbn(10)
	fmt.Println("fbn = ", data)
}

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	if n == 0 {
		return fbnSlice
	}
	for i := range fbnSlice {
		if i == 0 || i == 1 {
			fbnSlice[i] = 1
		} else {
			fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
		}
	}
	return fbnSlice
}
