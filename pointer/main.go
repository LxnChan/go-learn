package main

import "fmt"

func main() {
	a := "666"
	b := &a
	var c int = 996
	fmt.Println(b, *b)
	modify1(c)
	fmt.Println(c)
	modify2(&c)
	fmt.Println(c)
}

func modify1(x int) {
	x = 955
	return
}

func modify2(y *int) {
	*y = 1055
	return
}
