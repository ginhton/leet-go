package main

import (
	"fmt"
)

const MAX int = 3

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("address of a is %x\n", &a)
	fmt.Printf("ponter adderss; %x\n", ip)
	fmt.Printf("value of *ip: %d\n", *ip)

	a := []int{10, 20, 30}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d]=%d\n", i, *ptr[i])
	}

	var t int = 3
	var ptr *int = &t
	var pptr **int = &ptr
	var ppptr ***int = &pptr

	fmt.Println(t, *ptr, **pptr, ***ppptr)

	var x int = 3
	var y int = 4

	swap(&x, &y)

}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
