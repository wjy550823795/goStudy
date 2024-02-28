package main

import "fmt"

func changeValue(p int) {
	p = 10
}
func changeValue2(p *int) {
	*p = 10
}
func main() {
	//var a int = 1
	//changeValue(a)
	//fmt.Print("a=", a) //结果为1

	//var a int = 1
	//changeValue2(&a)
	//fmt.Print("a=", a) //结果为10

	var a int = 10
	var b int = 20
	//swap(a, b)
	swap2(&a, &b)
	fmt.Println("a=", a, "b=", b)
}

func swap(a int, b int) {
	var temp int = a
	a = b
	b = temp
}

func swap2(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
