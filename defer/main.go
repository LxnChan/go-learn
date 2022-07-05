package main

import "fmt"

//defer：延迟执行，先被defer的语句最后被执行
//等整个函数即将结束但还未结束时逆序执行
func main() {
	fmt.Println("Start...") //
	defer fmt.Println("1")  // 执行至此时将defer语句依次沉入池子底部，待函数内全部语句执行结束后再倒序执行（从池子最上端开始执行）
	defer fmt.Println("2")  // defer适用于执行后清理和记录时间、日志等情况。
	defer fmt.Println("3")
	fmt.Println("end")
}
