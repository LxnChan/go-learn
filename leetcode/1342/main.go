package main

import "fmt"

func main() {
	a := numberOfSteps(14)
	fmt.Println(a)
}

func numberOfSteps(num int) int {
	var numNum int
	for i := 0; i <= 100; i++ {
		if num == 0 {
			return numNum
			break
		} else if num%2 == 0 {
			num = num / 2
			numNum = numNum + 1
		} else {
			num = num - 1
			numNum = numNum + 1
		}
	}
	return numNum
}
