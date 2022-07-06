package main

import (
	"fmt"
	"strconv"
)

//import "strings"

func main() {
	outputResult := fizzBuzz(999)
	fmt.Println(outputResult)
}

func fizzBuzz(n int) []string {
	arrResult := make([]string, 0, 100000)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arrResult = append(arrResult, "FizzBuzz")
		} else if i%3 == 0 {
			arrResult = append(arrResult, "Fizz")
		} else if i%5 == 0 {
			arrResult = append(arrResult, "Buzz")
		} else {
			arrResult = append(arrResult, strconv.Itoa(i))
		}
	}
	return arrResult
}
