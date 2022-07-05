//define package main
//注释不能嵌套使用
package main

//import packages
import (
	"fmt"
)

//func-定义一个函数
//main-必须包含，是程序的入口函数
//花括号的开始不能单独一行
//逻辑必须在函数内部
func main() {
	//引用fmt包，引用里面的函数，在末尾自动添加ln（换行）
	fmt.Println("Hello.")
	var name string
	var age int
	//批量注册变量
	var (
		newsletter int
		c          bool
	)
	var d bool = true
	name = "I'm NiuBi!"
	age = 64
	c = false
	newsletter = 111
	fmt.Println(name, age, c, newsletter, d)

}
