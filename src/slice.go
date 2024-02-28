package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type is %T\n", myArray)
	printArraySlice(myArray)
	fmt.Println("==============")
	for _, val := range myArray {
		fmt.Println("val=", val)
	}
}
func printArraySlice(myArray []int) {
	for _, val := range myArray {
		fmt.Println("val=", val)
	}
	myArray[0] = 100
}
