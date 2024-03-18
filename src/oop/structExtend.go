package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

/*==============================================*/
// 子类
type SuperMan struct {
	Human //把父类对象注入到子类属性中,和java差异很大
	level int
}

// 子类重写
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Flow() {
	fmt.Println("SuperMan.Flow()...")
}
func (this SuperMan) print() {
	fmt.Println("superMan.name=", this.name)
	fmt.Println("superMan.sex=", this.sex)
	fmt.Println("superMan.level=", this.level)
}

func main() {
	h := Human{"zhang3", "haha"}
	h.Eat()
	h.Walk()
	fmt.Println("===================")

	//s := SuperMan{Human{"zhangsan", "男"}, 100}
	//两种定义方式
	var s SuperMan
	s.name = "li4"
	s.sex = "女"
	s.level = 99
	s.Walk() //父类方法
	s.Eat()  //子类重写方法
	s.Flow() //子类的新方法
	s.print()

}
