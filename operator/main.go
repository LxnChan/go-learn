package main

import "fmt"

//运算符
func main() {
	//算术运算符
	var a int = 10
	var b int = 90
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	//关系运算符
	fmt.Println(a != b)
	fmt.Println(a < b)

	//逻辑运算符
	fmt.Println(!(a < b))
	fmt.Println((a == b) && (a > b))

	//赋值运算符
	var c = 255
	c += 5 // => c=c+5
	fmt.Println(c)

	var d = 5
	// 5除以2，结果为1
	fmt.Println(d % 2)
}
