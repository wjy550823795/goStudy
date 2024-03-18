package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	getType() string  //获取动物的种类
}

// 具体的类
type Cat struct {
	color string
}

type Dog struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) getType() string {
	return "Cat"
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) getType() string {
	return "Dog"
}

func show(an AnimalIF) {
	an.Sleep() //多态
	fmt.Println("color = ", an.GetColor())
	fmt.Println("type = ", an.getType())
}

func main() {
	//var animal AnimalIF = &Cat{"white"}
	//animal.Sleep()
	Dog := Dog{"yellow"}
	Cat := Cat{"white"}
	show(&Dog)
	show(&Cat)

}
