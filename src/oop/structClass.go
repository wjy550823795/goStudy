package main

import "fmt"

// 如果类名首字母大写，表示其他包也能访问
type Hero struct {
	//如果类属性首字母大写，表示该属性可对外访问
	Name  string
	Ad    int
	level string
}

/*func (this Hero) show() {
	fmt.Println("Name=", this.Name, "Ad=", this.Ad, "Level=", this.Level)
}
func (this Hero) GetName() string {
	return this.Name
}
func (this Hero) SetName(newName string) {
	//当前this是一个副本（拷贝）
	this.Name = newName
}*/

// 相当于private
func (this *Hero) show() {
	fmt.Println("Name=", this.Name, "Ad=", this.Ad, "Level=", this.level)
}

// 相当于public
func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(newName string) {
	//当前this是一个副本（拷贝）
	this.Name = newName
}

func main() {

	hero := Hero{"wang", 1, "max"}
	hero.show()
	hero.SetName("jun")
	hero.show()
}
