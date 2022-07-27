package main

import "fmt"

//  中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1)
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = doSomething(s, i, i, res)
		res = doSomething(s, i, i+1, res) //偶数情况
	}
	return res
}

func doSomething(s string,start,end int,res int)int{
	for start >=0 && end < len(s) && s[start] == s[end]{
		res++
		end++ //中心扩散，所以右边右移，
		start-- //左边左移
	}
	return res
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
