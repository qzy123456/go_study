package main

import "fmt"

//找到缺失的第一个正整数
func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v] = v
	}
	for index := 1; index < len(nums)+1; index++ {
		if _, ok := numMap[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}

func main() {
   nums := []int{7,8,9,11,12}
   fmt.Println(firstMissingPositive(nums))
}
