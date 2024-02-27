package main

import "fmt"

func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a, "b=", b)
	c := 100
	return c, 777
}

func foo1(a string, b int) int {
	fmt.Println("a=", a, "b=", b)
	c := 100
	return c
}
func main() {
	c := foo1("123", 2)
	fmt.Println("c=", c)

	res1, res2 := foo2("haha", 666)
	fmt.Println("res1=", res1, "res2=", res2)
}
