package main

import (
	"fmt"
	"strconv"
	"strings"
)

//给定一个经过编码的字符串，返回它解码后的字符串。
//
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//示例 1：
//
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//示例 2：
//
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//示例 3：
//
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//示例 4：
//
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
var (
	src string
	ptr int
)
//递归
func decodeString(s string) string {
	src = s
	ptr = 0
	return getString()
}

func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits()
		ptr++
		str := getString()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigits() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}

//堆栈
func decodeString1(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		by := s[i]
		switch by {
		case ']':
			fmt.Println("遇到右中括号】，", string(stack))
			//字符串临时变量
			cnt := make([]byte, 0, len(s))
			//出栈，如果不是左括号，那么就是字符串，就放到字符串临时变量里面
			for stack[len(stack)-1] != '[' {
				cnt = append(cnt, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			fmt.Println("删除前-》", string(stack))
			//删除 '['，下面开始处理拼接
			stack = stack[:len(stack)-1]
			fmt.Println("删除后-》", string(stack))
			//临时数字变量
			count := ""
			//如果是数字，拼接数字
			for len(stack) != 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				count = count + string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//字符串转数字，求出具体数量
			num := 0
			for i := len(count) - 1; i >= 0; i-- {
				v, _ := strconv.Atoi(string(count[i]))
				num = num*10 + v
			}
			fmt.Println("待处理", num, string(cnt))
			//字符拼接，数字有多大，就写多少个字符串
			for i := 0; i < num; i++ {
				for j := len(cnt) - 1; j >= 0; j-- {
					stack = append(stack, cnt[j])
				}
			}
		default: //压栈
			stack = append(stack, by)
		}
	}
	return string(stack)
}

func main() {
	s := "10[ac]2[bc]"
	fmt.Println(decodeString1(s))
}
