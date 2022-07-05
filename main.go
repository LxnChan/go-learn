package main

import "fmt"

func main() {
	fmt.Println("Hello world.")

	//函数可以作为变量
	var gg = devTools(6, 8)
	fmt.Println(gg)
}

func devTools(a, b int) (c int) {
	//局部变量只在函数内部有效
	c = a * b
	return
}
