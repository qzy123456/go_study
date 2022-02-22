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

func main() {
	haystack := "aabaa"
	needle := "ba"
	fmt.Println(strStr(haystack, needle))
	fmt.Println(strStr2(haystack, needle))
}
