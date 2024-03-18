package main

import "fmt"

// 此类略过 傻逼语法糖
// const定义枚举类型
const (
	//可以在const()添加一个关键字iota，每行iota都会累加1，第一行默认值为0
	BEIJING = iota //iota=0
	SHANGHAI
	TIANJIN
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	//length = 100 //常量不允许修改
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(TIANJIN)
}
