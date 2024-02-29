package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	//添加 类似put
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printCityMap(cityMap)
	fmt.Println("===============")
	//删除
	delete(cityMap, "Japan")

	//修改
	cityMap["USA"] = "DC"
	changeCityMap(cityMap)
	printCityMap(cityMap)

}
func changeCityMap(cityMap map[string]string) {
	//引用传递
	cityMap["England"] = "London"
}

// 遍历
func printCityMap(cityMap map[string]string) {
	//
	for key, val := range cityMap {
		fmt.Println("key=", key, "val=", val)
	}
}
