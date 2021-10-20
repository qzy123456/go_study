package main

import "fmt"

//给定一个没有重复数字的序列，返回其所有可能的全排列
func permute(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	/*以下为三重循环*/
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < l; k++ {
				/*确保 i 、j 、k 三位互不相同*/
				if nums[i] != nums[k] && nums[i] != nums[j] && nums[j] != nums[k] {
					fmt.Println( "i =", i, "j =", j, "k =", k)
					res = append(res,[]int{nums[i],nums[j],nums[k]})
				}
			}
		}
	}
	return  res
}

//深度搜索DFS
func permute2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
	fmt.Println(permute2(nums))
}