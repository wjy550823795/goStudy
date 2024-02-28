package main

import "fmt"

func main() {

	//长度len为3
	//slice := []int{1, 2, 3}

	//没有分配空间 这时候赋值会报错
	var slice []int
	slice = make([]int, 3) //默认值为0 长度为3 赋值可以成功

	//var slice []int = make([]int, 3)
	//slice := make([]int, 3)

	//判断一个slice是否为空
	if slice == nil {
		fmt.Println("slice是空切片")
	} else {
		fmt.Println("slice 有空间")
	}

	fmt.Printf("len = %d,slice = %v\n", len(slice), slice)
}
