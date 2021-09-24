package main

import (
	"fmt"
)

func main() {
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance3 = [5]float32{1:2.0, 3:7.0}

	var n [10]int
	var i,j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Println("the first element is: ")
	fmt.Println(values[0][0])

	array_init := [3][4]int{
		{0, 1, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11}} // 最后的右大括号不能单独一行，除非前面一行用逗号结尾

	sites := [2][2]string{}

	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	fmt.Println(sites)
}
