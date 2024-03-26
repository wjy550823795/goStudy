package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
*
不退出goroutine
B
B.defer
A
A.defer
退出goroutine
B.defer
A.defer
*/
func main() {
	//用go 创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//退出当前的goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		defer fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
