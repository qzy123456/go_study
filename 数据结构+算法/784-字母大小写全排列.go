package main

import "fmt"

//784. 字母大小写全排列
//给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
//返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
//示例 1：
//
//输入：s = "a1b2"
//输出：["a1b2", "a1B2", "A1b2", "A1B2"]
//示例 2:
//
//输入: s = "3z4"
//输出: ["3z4","3Z4"]

func letterCasePermutation(s string) []string {
	var ans []string
	backtrace([]byte(s), 0,  &ans)
	return ans
}

func backtrace(str []byte, start int, ans *[]string) {
	*ans = append(*ans, string(str))
	for i:=start; i<len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] -= 32
			backtrace(str, i+1, ans)
			str[i] += 32
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] += 32
			backtrace(str, i+1, ans)
			str[i] -= 32
		}
	}
}
func main() {
	fmt.Println(letterCasePermutation("a1b1"))
}