package main

import (
	"fmt"
	"strconv"
)

//给定两个字符串形式的非负整数 num1 num2 计算他们的和

func addStrings(num1,num2 string) int  {
	add := 0  //进位标志
	ans :="" //最终字符串
	result := 0
	var i  = len(num1) -1 //末尾
	var j  = len(num2) -1 //末尾
	for i>=0 || j>= 0 || add != 0 {
		var x,y int
		if i >0 {
			x = int(num1[i] - '0')
		}
		if j>0 {
			y = int(num2[j] - '0')
		}
		result = x +y + add //计算num【i】+ num【j】+add的合
		ans = strconv.Itoa(result%10) + ans  //每次拼接一位，因为result可能大于10，所以对10取余
		add = result / 10
		i--
		j--
	}

	return result
}

func main() {
	num1 := "123"
	num2 := "987"
	fmt.Println(addStrings(num1,num2))
}
