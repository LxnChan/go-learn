package main

import "fmt"

func main() {
	println(plusValue(1, 3))
	fmt.Println(plusValue2(9, 3))
	var _, outputSlice2 int = calcValue(1, 3, 5)
	fmt.Println(outputSlice2, "/and then/")
}

func plusValue(a, b int) int {
	return a + b
}

//如果没有定义返回值就必须在函数内部新定义返回值
// func 函数名 (输入值)(返回值){函数体 return}
func plusValue2(c, d int) (e int) {
	e = c * d
	return e
}

//多返回值
func calcValue(input1, input2, input3 int) (output1, output2 int) {
	output1 = (input1 * input2) / input3
	output2 = input3*input1 + input2
	return
}
