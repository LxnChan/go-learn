package main

import "fmt"

func main() {
	var arr0 = []int{1, 3, 5, 7, 9}
	runningSum(arr0)
}

func runningSum(nums []int) []int {
	var arrLens = len(nums)
	//fmt.Println(arrLens)
	for i := 1; i < arrLens; i++ {
		nums[i] += nums[i-1]
	}
	fmt.Println(nums)
	return nums
}
