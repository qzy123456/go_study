package main

import (
	"fmt"
	"strings"
)

//验证是不是回文字符串
//“回文串”是一个正读和反读都一样的字符串，比如“level”或者“noon”等等就是回文串。
//对于字符串中可能存在的其他字符，可以通过正则替换，但是正则替换会增加程序运行复杂度，
// 下面给出的是在判断过程中忽略其他字符：
func isPalindromee(s string) bool {
	if s == "" {
		return false
	}
	s = strings.ToLower(s)
	if len(s) == 2{
		return s[0] == s[1]
	}

	left := 0
	right := len(s) - 1
	for left < right{
        //左边
        if !(s[left] >= 'a' && s[left] <= 'z') || (s[left] > 0 && s[left] <= 9){
        	left++
        	continue
		}
        //右边
		if !(s[right] >= 'a' && s[right] <= 'z') || (s[right] > 0 && s[right] <= 9){
			right--
			continue
		}
        if s[left] != s[right]{
        	return  false
		}
        left++
        right--
	}

	return  true
}



func main() {
	str :=  "ac! ca "
	fmt.Println(isPalindromee(str))
}
