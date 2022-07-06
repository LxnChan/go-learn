package main

import "fmt"

func main() {
	// 定义切片
	var sliceAtr = []int{66, 4, 55, 11, 100, 250}
	fmt.Println(sliceAtr[3])

	// 从数组获得切片
	var arrNtr = [...]string{"Natsuki", "Torari", "Klir"}
	sliceNtr := arrNtr[0:3]
	fmt.Println(sliceNtr)

	// make函数构造切片
	var sliceMrr = make([]int, 5, 10)
	fmt.Println(sliceMrr)

	// 获取切片容量
	var arrGround = [...]int{15, 23, 48, 19, 75, 16}
	sliceGround := arrGround[0:6]
	fmt.Println(sliceGround)
	fmt.Println(cap(sliceGround))
}
