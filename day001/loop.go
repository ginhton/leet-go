package main

import "fmt"

func main() {
sum := 0
for i := 0; i <= 10; i++ {
sum += 1;
}
fmt.Println(sum)

strings := []string{"google", "runoob"}
for i, s := range strings {
fmt.Println(i, s);
}

numbers := [6]int{1,2,3,5}
for i,x := range numbers {
fmt.Printf("number %d is at the index %d\n", x, i)
}


}
