package main

import (
	"goStudy/src/lib1"
	//执行init 匿名导入 无法使用当前包的方法，但会执行init方法
	//_ "goStudy/src/lib2"
	"goStudy/src/lib2"
	//给包设置别名
	mylib2 "goStudy/src/lib2"
	//将lib2包里的所有的方法导入当前类
	. "goStudy/src/lib2"
)

func main() {
	//执行接口
	lib1.Lib1Test()
	lib2.Lib2Test()
	mylib2.Lib2Test()
	Lib2Test()
}
