package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("muFunc is called...")
	fmt.Println(arg)
}

type BookBook struct {
	auth string
}

func main() {
	book := BookBook{"Golang"}
	myFunc(book)
	myFunc(100)

}
