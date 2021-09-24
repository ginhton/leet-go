package main

import "fmt"

func main() {


     var a int = 100;

     if a < 20 {
     fmt.Printf("a is less than 20\n")
     } else if a == 20 {
     fmt.Printf("a is equal to 20\n")
     } else {
     fmt.Printf("a is greater than 20\n")
     }
     fmt.Printf("a is %d\n", a);

     var grade string = "B"
     var marks int = 90

     switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 50,60,70: grade = "C"
        default: grade = "D"
     }

    switch {
        case grade=="A" :
        fmt.Println("brilliant!\n")
        case grade=="B", grade=="C" :
        fmt.Println("good!\n")
        case grade=="D" :
        fmt.Println("pass")
        case grade=="F" :
        fmt.Println("not pass!\n")
        default :
        fmt.Println("oh my god!\n")
    }
    fmt.Printf("Your grade is %s\n", grade)

    var x interface {}

    switch i := x.(type) {
    case nil:
    fmt.Printf("type of x is %T", i)
    case int:
    fmt.Printf("x is int")
    case float64:
    fmt.Println("x is float64")
    case func(int) float64:
    fmt.Println("x is func(int)")
    case bool, string:
    fmt.Printf("x is bool or string")
    default:
    fmt.Printt("unknown type")
    }

    switch {
    case false:
    fmt.Println("1. case false")
    fallthrough
    case true:
    fmt.Println("2. case true")
    fallthrough
    case false:
    fmt.Println("3. case false")
    fallthrough
    case true:
    fmt.Println("4. case true")
    case false:
    fmt.Println("5. case false")
    fallthrough
    default:
    fmt.Println("6. default")
    }

    select {
    case communication clause :
    statement(s);
    case communication clause :
    statement(s);
    /* ... */
    default :
    statement(s);
    }

    var c1, c2, c3 chan int
    var i1, i2 int

    select {
    case i1 = <-c1:
    fmt.Printf("received ", i1, " from c1\n")
    case c2 <- i2:
    fmt.Printf("sent ", i2 " to c2\n")
    case i3, ok = (<-c3): // same as i3, ok := <-c3
        if ok {
        fmt.Println("received ", i3, " from c3\n")
        } else {
        fmt.Printf("c3 is closed\n")
        }
    default :
    fmt.Printf("no communication\n")
    }

}
