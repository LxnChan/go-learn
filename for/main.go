package main

import "fmt"

func main() {
	//完整格式
	var a int
	for a = 102; a < 300; a++ {
		if a%2-1 == 0 {
			fmt.Println(a)
		}
	}

	//省略格式
	var b int
	for b < 100 {
		b++
		if b%2-1 == 0 {
			fmt.Println(b)
		}
	}

	//break跳出循环
	var c int
	for {
		c++
		fmt.Println(c)
		if c == 100 {
			break
		}
	}

	//continue继续循环
	//var d int
	//for {
	//	if d == 4 {
	//		continue
	//	}
	//	d++
	//	fmt.Println(d)
	//}

	//switch语句
	var e = 5
	switch e {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 4, 6, 8, 10:
		fmt.Println("666")
	//只能有一个default语句
	default:
		fmt.Println("0")
	}

	//case中可以写表达式，switch后面的参数留空
	f := 18
	switch {
	case f < 16:
		fmt.Println("Young")
	default:
		fmt.Println("old!")
	}
}
