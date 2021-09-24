package main

import (
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Golang tutorial", "www.runoob.com", "tutorial", 65222})

	fmt.Println(Books{title: "Golang tutorial", author: "www.runoob.com", subject: "tutorial", book_id:65222})

	// fields without specified value will be 0 or "" or nil
	fmt.Println(Books{title: "Golang tutorial", author :"www.runoob.com"})

}
