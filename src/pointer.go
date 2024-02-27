package main

import "fmt"

func changeValue(p int) {
	p = 10
}
func main() {
	//var a int = 1
	//changeValue(a)
	//fmt.Print("a=", a) //结果为1

	var a int = 1
	changeValue2(&a)
	fmt.Print("a=", a) //结果为10
}

func changeValue2(p *int) {
	*p = 10
}
