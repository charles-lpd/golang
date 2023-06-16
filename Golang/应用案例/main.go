package main

import "fmt"

func main() {
	// 1) 统计 3 个班成绩情况，每个班有5名同学，
	// 求出各个班的平均分，和所有班级学生的平均分[学生的成绩从键盘输入]
	classNum := 3
	quantity := 5
	var classes1 float64 = getSum(quantity, "一")
	fmt.Println("第一个班级的平均分是 \n", classes1/float64(quantity))
	var classes2 float64 = getSum(quantity, "二")
	fmt.Println("第一个班级的平均分是\n", classes2/float64(quantity))
	var classes3 float64 = getSum(quantity, "二")
	fmt.Println("第一个班级的平均分是\n", classes3/float64(quantity))

	total := classes1 + classes2 + classes3

	fmt.Println("所有班级学生的总平均分是", total/float64(quantity*classNum))
}
func getSum(num int, str string) float64 {
	sum := 0.0
	for i := 1; i <= num; i++ {
		score := 0.0
		fmt.Printf("请输入第%s班，第%d 个学生的成绩 \n", str, i)
		fmt.Scanln(&score)
		sum += score
	}
	return sum
}
