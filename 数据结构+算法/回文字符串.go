package main

import "fmt"

//方法1
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}
//方法2 暴力解法
func longestPalindrome2(s string) string {
	var res string = string(s[0])
	for i:=0;i<len(s)-1;i++{
		for j:=i+1;j<len(s);j++{
			if palindrome(s[i:j+1]) && j-i+1>len(res){
				res = s[i:j+1]
			}
		}
	}

	return res
}

func palindrome(s string)bool{
	var n = len(s)
	if n==1{
		return true
	}
	for i:=0;i<n/2;i++{
		if s[i]!=s[n-i-1]{
			return false
		}
	}
	return true
}

func main() {

	str := "babad"
	fmt.Println(longestPalindrome(str))
	fmt.Println(longestPalindrome2(str))
}
