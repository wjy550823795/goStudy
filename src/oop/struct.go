package main

import "fmt"

// 声明一种数据类型是myint
type myint int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	//传递副本，不会改本main::book的值
	book.author = "jun"
}

func changeBookPoint(book *Book) {
	//传递副本，不会改本main::book的值
	book.author = "jun"
}

func main() {
	var a myint = 10
	fmt.Println("a=", a)
	fmt.Printf("type of a =%T\n", a)
	var book Book
	book.title = "Golang"
	book.author = "wang"
	fmt.Printf("%v\n", book)
	fmt.Println("=============")
	changeBook(book)
	fmt.Printf("%v\n", book)
	fmt.Println("=============")
	changeBookPoint(&book)
	fmt.Printf("%v\n", book)
}
