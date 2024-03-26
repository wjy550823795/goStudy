package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Gorountine:i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	//创建一个go程 去执行newTask()流程
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main Gorountine:i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
