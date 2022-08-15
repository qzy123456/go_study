package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
    reverseString := func(s []byte) {
        for i, j := 0, len(s)-1; i < j; {
            s[i], s[j] = s[j], s[i]
            i++
            j--
        }
    }
    sb := []byte(s)

    for i, j := 0, 0; i < len(sb); i = j {
        for j < len(sb) && sb[j] != ' ' {
            j++
        }
        reverseString(sb[i:j])
        j++ // 跳到下一个单词的头
    }
    return string(sb)
}
func reverseWords2(s string) string {
	s1:=strings.Split(s," ")
	for r:=range s1{
		s1[r]=string(reverse([]byte(s1[r])))
	}
	return strings.Join(s1," ")
}
func reverse(s []byte) []byte{
	f,e:=0,len(s)-1
	for f<e{
		s[f],s[e]=s[e],s[f]
		f,e=f+1,e-1
	}
	return s
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords2(s))
}