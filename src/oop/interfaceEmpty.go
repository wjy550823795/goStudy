package main

import "fmt"

// 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("muFunc is called...")
	//如何区分数据类型 提供断言机制
	value, ok := arg.(string) //这语法糖真垃圾 没有提示的
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type，val=", value)
	}
}

type BookBook struct {
	auth string
}

func main() {
	book := BookBook{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")

}
