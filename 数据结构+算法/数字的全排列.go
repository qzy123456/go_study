package main

import "fmt"

//给定一个没有重复数字的序列，返回其所有可能的全排列
//这种是不行的，因为测试用例会有数组为2个的情况 [1,2]
func permute(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	/*以下为三重循环*/
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < l; k++ {
				/*确保 i 、j 、k 三位互不相同*/
				if nums[i] != nums[k] && nums[i] != nums[j] && nums[j] != nums[k] {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}
var res [][]int
func permute2(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums,len(nums),[]int{})
	return res
}
func backTrack(nums []int,numsLen int,path []int)  {
	if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}
	for i:=0;i<numsLen;i++{
		cur:=nums[i]
		path = append(path,cur)
		nums = append(nums[:i],nums[i+1:]...)//直接使用切片
		backTrack(nums,len(nums),path)
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)//回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}

func permute3(nums []int) [][]int {
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

		for i := range nums {
			if used[i] {
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
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	fmt.Println(permute2(nums))
	fmt.Println(permute3(nums))
}
