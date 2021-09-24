package main

import "fmt"

type cb func(int) int

func max(num1, num2 int) int {
	var result int

	if (num1 < num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x;
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printlf("the max value is %d\n", ret)

	c, d := swap("Google", "Runoob")
	fmt.Println(c, d)


	testCallback(1, callBack)
	testCallback(2, func(x int) int {
		fmt.Printf("I'm callback, x: %d\n", x)
		return x
	})
}

func testCallback(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Println("I'm callback, x: %d\n", x)
	return x;
}
