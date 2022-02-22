package main

import "fmt"

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	if last < 0 {
		return 0
	}
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	return last - first
}

func lengthOfLastWord2(s string) int {
	last := len(s) - 1
	res := 0
	//最后一个单词，就倒叙查找
	for i:=last;i>0 ;i--  {
		if s[i] == ' '{ //判断如果是空格
			if res > 0{ //如果是单词之后才碰到的空格
				return res
			}
			continue
		}
		res++
	}
	return res
}

func main() {
	words := "hello word1 "
	fmt.Println(lengthOfLastWord(words))
	fmt.Println(lengthOfLastWord2(words))
}