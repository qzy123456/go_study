package main

import "fmt"
//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//输入: s = "abca"
//输出: true
//解释: 你可以删除c字符。
func validPalindrome(s string) bool {
	start := 0
	end := len(s)-1
	for start <end{
		if s[start] == s[end]{
			start++
			end--
		}else{
			//如果不满足，跳过
			return doSomething(s,start+1,end) || doSomething(s,start,end-1)
		}
	}
	return true
}
func doSomething(s string,start,end int)bool{
	for start < end{
		if s[start] == s[end]{
			start++
			end--
		}else{
			return false
		}
	}
	return  true
}

func main()  {
	fmt.Println(validPalindrome("aba"))
}
