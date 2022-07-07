package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x<0{
		return false
	}

	var result int
	for i:=x; i!=0; i=i/10{
		fmt.Println("i",i)
		tmp := i%10
		fmt.Println("temp",tmp)
		result = result*10 + tmp
		fmt.Println("result",result)
	}

	return result==x
}
// 解法二 数字转字符串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome1(1221))
}
