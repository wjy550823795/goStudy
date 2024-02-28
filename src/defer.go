package main

import "fmt"

func main() {
	defer fmt.Println("main end1111") //类似final
	defer fmt.Println("main end2222") //stack 先入后出
	fmt.Println("main hello 1")
	fmt.Println("main hello 2")
}
