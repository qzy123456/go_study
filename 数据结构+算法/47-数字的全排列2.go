package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列，返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑
	generatePermutation47(nums, 0, p, &res, &used)
	return res
}

func generatePermutation47(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 这里是去重的关键逻辑
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation47(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}
//46题的解法
func permuteUnique2(nums []int) [][]int {
	var res, path = make([][]int, 0), make([]int, 0)
	var used = make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i:=0;i<len(nums);i++ {
			if used[i] {
				continue
			}
			//关键在这一句
			if i != 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			// 回溯的过程中，将当前的节点从 path 中删除
			path = path[:len(path) - 1]
			used[i] = false
		}
	}
	dfs()
	return res
}

func main() {
	nums := []int{1,1,3}
	fmt.Println(permuteUnique(nums))
	fmt.Println(permuteUnique2(nums))
}