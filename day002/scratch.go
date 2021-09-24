package main

import (
	"fmt"
)

func main() {
	var m map[int]int
	m = make(map[int]int)
	m[1] = 3;
	i,ok := m[2]
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("not exist")
	}
}
