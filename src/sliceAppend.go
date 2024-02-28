package main

import "fmt"

func main() {
	//长度为3 容量为5
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d,cap=%d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len = %d,cap=%d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len = %d,cap=%d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3) //cap*2
	fmt.Printf("len = %d,cap=%d, slice = %v\n", len(numbers), cap(numbers), numbers)
}
