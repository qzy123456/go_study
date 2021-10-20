package main

import (
	"fmt"
	"sort"
)

//根据一个数组，找到相加满足target的1个或者多个数字
//1：可以无限使用
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化
			break
		}
		c = append(c, nums[i])
		findcombinationSum(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}
////根据一个数组，找到相加满足target的1个或者多个数字
////2：每次遍历数字只能用一次
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	sort.Ints(candidates) // 这里是去重的关键逻辑
	findcombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findcombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
func main() {
	nums := []int{1,3,2,4}
	fmt.Println(combinationSum(nums,5))
	fmt.Println(combinationSum2(nums,5))
}
