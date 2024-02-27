package lib1

import "fmt"

// 首字母大写表示对外开放接口
func Lib1Test() {
	fmt.Println("Lib1Test() ...")
}
func init() {
	fmt.Println("lib1.init() ...")
}
