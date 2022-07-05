package main

import (
	"fmt"
)

//数组
func main() {
	//定义数组
	//  数组名    [数量] 类型 {元素1,...,元素n-1}
	var arrCity = [4]string{"beijing", "shanghai", "huangzhou", "shenzhen"}
	fmt.Println(arrCity)

	//编译器推导数组数量
	var arrTransfer = [...]string{"1", "2", "3"}
	//依据角标提取数组中的数据
	fmt.Println(arrTransfer[2], "niubia")
	//数组长度：len(数组)
	fmt.Println("城市数量为", len(arrCity))

	//索引值方式
	var arrPeople = [...]string{1: "小王", 3: "18", 7: "女"}
	fmt.Println("姓名为", arrPeople[1], "，年龄为", arrPeople[3], "，性别为", arrPeople[7])

	//遍历
	var arrAge = [...]int{12, 13, 14, 15, 16, 17, 18, 19, 50, 49, 50, 61, 22, 33, 55}
	for i := 0; i <= len(arrAge)-1; i++ {
		if arrAge[i] >= 18 {
			fmt.Println(arrAge[i])
		} else {
			fmt.Println("太年轻不予显示")
		}
	}

	fmt.Println("----------------")

	//多维数组
	//数组嵌套数组，注意元素数量，里面嵌套几层就写几个
	//注意只有第一层可以使用[...]让编译器推导
	var arrRobot = [...][4]string{
		{"1a96", "2a13", "3a16", "4a19"},
		{"1a96b", "2a13b", "3a16b", "4a19b"},
	}
	fmt.Println(arrRobot[1][2])
	//多维遍历
	for i := 0; i <= len(arrRobot)-1; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Println(j)
			fmt.Println(arrRobot[i][j])
		}
	}

	//遍历求和
	var arrNum = [...]int{1, 3, 5, 7, 9}
	var numSi int
	for i := 0; i <= len(arrNum)-1; i++ {
		numSi = numSi + arrNum[i]
	}
	fmt.Println(numSi)

	fmt.Println("--------")

	//找指定下标
	var arrNum2 = [...]int{1, 3, 5, 7, 9}
	//var numSig2 int
	for i := 0; i <= len(arrNum2) && i+1 < len(arrNum2); i++ {
		if arrNum2[i]+arrNum2[i+1] == 10 {
			fmt.Println(i)
		}
	}
}
