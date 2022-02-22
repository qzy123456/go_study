package main

import (
	"fmt"
	"reflect"
)

// 用两个字典分别把两个字符串的字符出现个数保存起来
func isAnagram1(s string, t string) bool {
	var sMap = make(map[rune]int)
	var tMap = make(map[rune]int)

	for _, c := range s {
		sMap[c] = sMap[c] + 1
	}
	for _, c := range t {
		tMap[c] = tMap[c] + 1
	}
	return reflect.DeepEqual(sMap, tMap)
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram1(s,t))
}
