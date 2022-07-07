package main

import (
	"fmt"
	"sort"
)

//双指针  和3数之和一样  把index+j看成第三个数 n^4转为n^3
func fourSum(nums []int, target int) [][]int {
	//开始，结束，长度，和，返回值
	start, end, len, sum, res := 0, 0, len(nums), 0, [][]int{}
	sort.Ints(nums)
	//1，-1，0，0
	for j := 0; j < len; j++ {
		//第一个数
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}
		for index := j + 1; index < len; index++ {
			//第二个数
			if index > j+1 && nums[index] == nums[index-1] {
				continue
			}
			//index+1=2，len-1=3
			start, end = index+1, len-1 //第三个数  最后一个数
			for start < end {
				//sum = 0+1+2+3
				sum = nums[start] + nums[index] + nums[end] + nums[j]
				if sum < target { //小于
					start++
				} else if sum > target {  //大于
					end--
				} else { //刚好等于
					res = append(res, []int{nums[j], nums[index], nums[start], nums[end]})
					for start < end && nums[start] == nums[start+1] {
						start++
					}
					for start < end && nums[end] == nums[end-1] {
						end--
					}
					start++
					end--
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, -1, 0, 0}
	fmt.Println(fourSum(nums, 0))
}
