package main

import (
	"fmt"
	"time"
)

/*
*
有缓存的channel
*/
func main() {
	c := make(chan int, 3) //带有缓冲的channel
	fmt.Println("len(c)=", len(c), ",cap(c)", cap(c))

	go func() {
		defer fmt.Println("子 goroutine结束")
		for i := 0; i < 3; i++ {
			c <- i //send
			fmt.Println("子go正在运行:,发送的元素:", i, "len(c)=", len(c), ",cap(c)", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c //接收
		fmt.Println("num=", num)
	}
	fmt.Println("main goroutine 结束")
}
