package main

import (
	"fmt"
)

func main() {
	// 声明map并初始化
	var mapCity = make(map[string]string)
	fmt.Println(mapCity)
	mapCity["Beijing"] = "001"
	mapCity["Shanghai"] = "002"
	fmt.Println(mapCity)
	// 取出键值
	fmt.Println(mapCity["Shanghai"])
	//判断键值是否存在
	val, ok := mapCity["ShenYang"]
	fmt.Println(ok, val)
	// 删除键值
	//delete(mapCity, "Shanghai")
	//fmt.Println(mapCity)

	for k, v := range mapCity {
		fmt.Println(k, v)
	}
}
