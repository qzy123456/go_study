package main

import (
	"fmt"
)

// 解法一
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
func strStr2(haystack string, needle string) int {
	if len(needle) == 0{
		return 0
	}
	for loc := 0;loc<=len(haystack) - len(needle) ;loc++  {
		if haystack[loc:len(needle) + loc] == needle{
			return loc
		}
	}
	return -1
}
func strStr3(haystack string, needle string)int  {
	// next数组
	next := make([]int, len(needle))
	// 构建next数组的函数
	createNext := func() {
		// 当模式串长度为0，就直接返回
		if len(needle) == 0 {
			return
		}
		// 默认next[0]=-1
		k, index := -1, 0
		next[0] = -1
		// 使用模式串倒数第二个字符计算最后一个字符的next值，所以这里是<len(needle)-1
		for index < len(needle)-1 {
			// 当字符匹配，直接加1
			if k == -1 || needle[index] == needle[k] {
				index++
				k++
				next[index] = k
			} else {
				// 不匹配就更新k的值
				k = next[k]
			}
		}
	}
	// 模式匹配函数
	kmp := func() int {
		// 先处理特殊情况的空字符串
		if len(haystack) == 0 && len(needle) != 0 {
			return -1
		}
		if len(needle) == 0 {
			return 0
		}
		// 当连个字符串都不为空
		for h, n := 0, 0; ; {
			if haystack[h] == needle[n] {
				// 如果匹配，两个字符串的下标都加1
				n++
				h++
				// 是否到了模式串的最后
				if n == len(needle) {
					return h - n
				}
				// 是否到了被匹配字符串你的最后
				if h == len(haystack) {
					return -1
				}
			} else {
				// 遇到不匹配的字符
				n = next[n]
				// 从第一个字符就不匹配，那么就要将连个字符串都向后移动一个位置，因为只有next[0]=-1
				// 其他最小的是next的值是0，表示需要对模式串重新从最开始匹配
				if n < 0 {
					n++
					h++
					if h == len(haystack) {
						return -1
					}
				}
			}
		}
	}
	createNext()
	return kmp()
}

func main() {
	haystack := "abcde"
	needle := "de"
	fmt.Println(strStr(haystack, needle))
	fmt.Println(strStr2(haystack, needle))
	fmt.Println(strStr3(haystack, needle))
}
