package main

import "fmt"

func main() {

	var a string
	a = "abcd" //pair<static type:string,value:"abcd">

	var allType interface{}
	allType = a //pair<type:string,value:"abcd">
	fmt.Println(allType)

	res, _ := allType.(string)
	fmt.Println(res)
}
