package main

import (
	"fmt"
	"sort"
)

var res = [][]int{}

func nSum(nums []int, k int, target int) [][]int {
	sort.Ints(nums)
	trace := []int{}
	//kSum(k,0,target,nums,trace)
	kSum2(k, 0, target, nums, trace)
	return res
}

//防止超时
func kSum(k, start, target int, nums, trace []int) {
	length := len(nums)
	if k > 2 {
		for i := start; i <= length-k; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			trace = append(trace, nums[i])
			kSum(k-1, i+1, target-nums[i], nums, trace)
			//用过的值要剔除
			trace = trace[:len(trace)-1]
		}
	} else if k == 2 {
		left, right := start, length-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > target {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				//把临时数组的值取出来
				tmp := make([]int, len(trace))
				copy(tmp, trace) //拷贝
				tmp = append(tmp, nums[left], nums[right])
				//拼接返回值
				res = append(res, tmp)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
}
func kSum2(k, start, target int, nums, trace []int) {
	length := len(nums)
	if k == 0 && target == 0 {
		//拼接返回值
		res = append(res, trace)
		return
	}
	for i := start; i <= length-k; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		trace = append(trace, nums[i])
		kSum(k-1, i+1, target-nums[i], nums, trace)
		//用过的值要剔除
		trace = trace[:len(trace)-1]
	}

}

func main() {
	nums := []int{1, -1, 0, 2, -1, -1}
	fmt.Println(nSum(nums, 3, 0))
}
