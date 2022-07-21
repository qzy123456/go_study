package main

import (
	"fmt"
)

// 用两个字典分别把两个字符串的字符出现个数保存起来
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sMap = make(map[rune]int)
	var tMap = make(map[rune]int)

	for _, c := range s {
		sMap[c] += 1
	}
	for _, c := range t {
		tMap[c] += 1
	}
	for i, v := range sMap {
		vv, ok := tMap[i]
		//查不到，或者数量不一样
		if !ok || vv != v {
			return false
		}
	}
	return true
	//return reflect.DeepEqual(sMap, tMap)
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram1(s, t))
}
