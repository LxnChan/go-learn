package main

import "fmt"

func main() {
	var sliceAtr = [][]int{{1, 5}, {7, 3}, {3, 5}}
	op := maximumWealth(sliceAtr)
	fmt.Println(op)
}

func maximumWealth(accounts [][]int) int {
	var bestValue int = 0
	var currValue int = 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			currValue = currValue + accounts[i][j]
			//fmt.Println(currValue)
		}
		if currValue > bestValue {
			bestValue = currValue
		}
		// 强制清零
		currValue = 0
	}
	return bestValue
}
