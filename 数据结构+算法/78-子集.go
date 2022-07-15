package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
func subsets(nums []int) [][]int {
	res := make([][]int,1) //先定义一个空集合，空集合也算一个子集
	sort.Ints(nums)
	for i := range nums {
		//fmt.Println(res)
		//[[]]
		//[[] [1]]
		//[[] [1] [2] [1 2]]
		for _, org := range res {
			//fmt.Println(len(org))
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			//fmt.Println(clone)
			res = append(res, clone)
		}
	}
	return res
}

// 解法2 dfs回溯
func subsets2(nums []int) [][]int {
	var res [][]int
	var path []int
	generateSubsets(nums, 0, path, &res)
	return res
}

func generateSubsets(nums []int, startIndex int, path []int, res *[][]int) {
	//空数组也算一个子集
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)
	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		generateSubsets(nums, i + 1, path, res)
		path = path[:len(path) - 1]
	}
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
	fmt.Println(subsets2(nums))
	//[[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}