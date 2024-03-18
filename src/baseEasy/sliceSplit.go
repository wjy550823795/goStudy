package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s1 := s[0:2]
	fmt.Println(s1)

	s1[0] = 100 //s1变 s也跟着变
	s1 = append(s1, 666)
	fmt.Println("s=", s)
	fmt.Println("s1=", s1)

	//copy 可以将底层数组的slice数组一起进行拷贝
	s2 := make([]int, 3)
	//将s中的值拷贝到s2中
	copy(s2, s)
	fmt.Println("copy s2:", s2)
	s2[1] = 222 //s2变 s不变
	fmt.Println("copy s2[1] = 222:", s2)
	fmt.Println(s)
}
