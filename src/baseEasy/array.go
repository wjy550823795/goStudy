package main

import "fmt"

func main() {
	//固定长度
	var myArray1 [10]int
	//有初始值
	myArray2 := [10]int{1, 2, 3, 4}
	//长度不同 类型也不同
	myArray3 := [4]int{10, 20, 30, 40}
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	for index, val := range myArray2 {
		fmt.Println("index = ", index, "val = ", val)
	}
	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println("==========")
	for index, val := range myArray3 {
		fmt.Println("index = ", index, "val = ", val)
	}
}

func printArray(myArray [4]int) {
	for index, val := range myArray {
		fmt.Println("index = ", index, "val = ", val)
	}
	myArray[0] = 100
}
