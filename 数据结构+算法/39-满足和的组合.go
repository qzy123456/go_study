package main

import (
	"fmt"
	"sort"
)

//根据一个数组，找到相加满足target的1个或者多个数字
//1：可以无限使用
func combinationSum(candidates []int, target int) [][]int {
	var track []int
	var res [][]int
	backtracking(0,0,target,candidates,track,&res)
	return res
}
func backtracking(startIndex,sum,target int,candidates,track []int,res *[][]int){
	//终止条件
	if sum==target{
		tmp:=make([]int,len(track))
		copy(tmp,track)//拷贝
		*res=append(*res,tmp)//放入结果集
		return
	}
	if sum > target{
		return
	}
	//回溯
	for i:=startIndex;i<len(candidates);i++{
		//更新路径集合和sum
		track=append(track,candidates[i])
		sum+=candidates[i]
		//递归
		backtracking(i,sum,target,candidates,track,res)
		//回溯
		track=track[:len(track)-1]
		sum-=candidates[i]
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


//官网答案
func combinationSum3(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {
	nums := []int{1,3,2,4}
	fmt.Println(combinationSum(nums,5))
	fmt.Println(combinationSum2(nums,5))
}
