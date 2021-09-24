package main

import "unsafe"

const (
      Unknown = 0
      Female = 1
      Male = 2
)

const (
      x = "abc"
      y = len(a)
      z = unsafe.Sizeof(x)
)

func main() {
     const (
           a = iota // 0
           b        // 1
           c        // 2
           d = "ha" // "ha"; iota+=1
           e        // "ha", same as before if unspecified; iota+=1
           f = 100  // 100; iota+=1
           g        // 100, same as previous value if unspecified; iota+=1;
           h = iota // 7
           i        // 8
)
     println(a, b, c, d, e, f, g, h, i)
     println(x, y, z)
}
