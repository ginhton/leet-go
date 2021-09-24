package main

import (
       "fmt"
)

func main() {
     var stockcode=123;
     var enddate="2020-12-31"
     var url="Code=%d&endDate=%s"
     var target_url=fmt.Sprintf(url,stockcode,enddate)
     fmt.Println(target_url)

     var a string = "Runoob"
     fmt.Println(a)

     var b, c int = 1, 2
     fmt.Println(b, c)

     // 声明一个变量并初始化
     var aa = "RUNOOB"
     fmt.Println(aa)

     // 没有初始化就为零值
     var bb int
     fmt.Println(bb)

     // uninitialized bool is false, string is ""
     // other nil
     var a1 *int
     var a2 []int
     var a3 map[string] int
     var a4 chan int
     var a5 func(string) int
     var a6 error // interface

     // use :=
     f := "Runoob"
     fmt.Println(f)
}
