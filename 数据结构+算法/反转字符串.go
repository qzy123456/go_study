package main

import "fmt"

//循环
func reverseString(s []byte)  {
	for i:=0;i<len(s)/2 ;i++  {
		s[i],s[len(s) - i -1] = s[len(s) - i -1],s[i]
	}
}
//双指针
func reverseString2(s []byte)  {
	var end = len(s) -1
	var start = 0
	for end >start{
		s[start] = s[end]
		end--
		start++
	}
}

func main() {
   var s  = []byte{'a','b','c','d','e'}
	reverseString(s)
	var s1  = []byte{'a','b','c','d','e'}
	reverseString2(s1)
   fmt.Println(string(s))
   fmt.Println(string(s1))
}
