package main

import "fmt"

// 形参返回
func foo3() (r1 int, r2 int) {
	//初始化默认值为0 r1,r2为foo3形参
	fmt.Println("r1=", r1, "r2=", r2)
	r1 = 100
	r2 = 200
	return
}

// 匿名返回
func foo2(a string, b int) (int, int) {
	c := 100
	return c, 777
}

func foo1(a string, b int) int {
	c := 100
	return c
}
func main() {
	c := foo1("123", 2)
	fmt.Println("foo1 c=", c)

	res1, res2 := foo2("haha", 666)
	fmt.Println("foo2 res1=", res1, "res2=", res2)

	res1, res2 = foo3()
	fmt.Println("foo3 res1=", res1, "res2=", res2)
}
