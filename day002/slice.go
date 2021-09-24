package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	numbers1 := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers1)

	fmt.Println("numbers1 == ", numbers1)
	fmt.Println("numbers1[1:4] == ", numbers1[1:4])
	fmt.Println("numbers1[:3] == ", numbers1[:3])
	fmt.Println("numbers1[4:] == ", numbers1[4:])

	numbers2 := make([]int, 0, 5)
	printSlice(numbers2)

	numbers3 := numbers1[:2]
	printSlice(numbers3);

	numbers4 := numbers1[2:5]
	printSlice(numbers4)

	var nums []int
	printSlice(nums)

	nums = append(nums, 0)
	printSlice(nums)

	nums = append(nums, 1)
	printSlice(nums)

	nums = append(nums, 2, 3, 4)
	printSlice(nums)

	nums1 = make([]int, len(nums), (cap(nums))*2)

	copy(nums1, nums)
	printSlice(nums1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(x), cap(x), x)
}
