package main

import "fmt"

func main() {
	records := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	dbPage := 1
	count := 10

	start := dbPage * count
	end := start + count

	// if end > len(records) {
	// 	end = len(records)
	// }
	if start > len(records) {
		records = make([]int, 0)
		fmt.Println(records)
	} else {
		fmt.Println(end)
		records = records[start:end]
		fmt.Println(records)
	}

}
