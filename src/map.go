package main

import "fmt"

func main() {
	//声明map类型
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap 是empty")
	}
	//需要先发分配空间
	myMap = make(map[string]string, 10)
	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "golang"

	fmt.Println(myMap) //乱序

	/*=====================*/
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "golang"
	fmt.Println(myMap2) //乱序

	/*======================*/
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"trhee": "golang",
	}
	fmt.Println(myMap3) //乱序
}
