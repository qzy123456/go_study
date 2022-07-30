package main

import "fmt"

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//示例 1:
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//输入: nums = [1], k = 1
//输出: [1]
func topKFrequent(nums []int, k int) []int {
	// 用一个map来存储数字以及出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	//fmt.Println(m) //map[0:2 1:1 3:1]
	// 定义存储最后结果的数组
	res := make([]int,k) //分配返回值的数组，默认k个，初始化为0
	for i := 0; i < k; i++ {
		tempMax := 0
		for k, num := range m {
			if num > tempMax {
				res[i] = k
				tempMax = num
			}
		}
		m[res[i]] = -1
		//fmt.Println(m)
	}
	return res
}

func main() {
	//题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
	a := []int{1,1,1,2,2,3}
	fmt.Println(topKFrequent(a, 2))
}
