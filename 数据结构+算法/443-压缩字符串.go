package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int{
	lenght := len(chars)
	current := 0
	for i:=0;i<lenght;  {
		start := i
		chars[current] = chars[i]
		for i+1<lenght && chars[i+1] == chars[i] {
			i++
		}
		if i-start+1 == 1 {
			i++
			current++
		}else {
				strLen := strconv.Itoa(i-start+1)
			for j:=0;j<len(strLen) ;j++  {
				current++
				chars[current] = strLen[j]
			}
			current++
			i++
		}
	}
	return current
}
func main() {
	chars := []byte{'a','a','b','b','c'}
	fmt.Println(compress(chars))
}
