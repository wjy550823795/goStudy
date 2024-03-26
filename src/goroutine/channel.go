package main

import (
	"fmt"
	//"time"
)

func main() {
	//定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine 正在运行。。。")
		c <- 666 //将666发送给c
	}()

	num := <-c //从c中接收数据,并赋值给num
	fmt.Println("num=", num)
	//如果c的数据不被读取出来，goroutine会一直呗阻塞，直至主线程结束
	//time.Sleep(1 * time.Second)
	fmt.Println("main goroutine 结束")
}
