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

// 解法二 滑动窗口，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++

		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
// 解法三 中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome3(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

func longestPalindrome4(s string) string {
	if s == "" {
		return ""
	}
	if palindrome(s) || len(s) < 2 {
		return s
	}

	max := 0
	rst := ""

	for i := 0; i < len(s); i++ {
		for j:=len(s);j>i+1;j--{
			tmp := s[i:j]
			//fmt.Println(tmp)
			if palindrome(tmp) {
				if max < len(tmp) {
					max = len(tmp)
					rst = tmp
				}
			}
		}
	}

	if rst == "" {
		return s[0:1]
	}
	return rst
}
func main() {

	str := "bababad"
	///fmt.Println(longestPalindrome(str))
	fmt.Println(longestPalindrome2(str))
	//fmt.Println(longestPalindrome1(str))
	//fmt.Println(longestPalindrome3(str))
}
