package main

import "fmt"

//if 语句
func main() {
	var nb = 255
	if nb == 255 {
		fmt.Println("nice")
	} else if nb <= 255 {
		fmt.Println("not nice")
	} else {
		fmt.Println("done!")
	}
	if nbhl := 999; nbhl < 5 {
		fmt.Println("not true")
	} else {
		fmt.Println("666")
	}
}
