package main

import (
	"fmt"
	"sort"
)

// 输入 nums =[10,9,2,5,3,7,101,18]
//最长递增子序列是 2,3,7,101 所以长度是4

//动态规划+sort.SearchInts()
func lengthOfLIS(nums []int) int {
	var dp []int
	for _, num := range nums {
		i := sort.SearchInts(dp, num) //min_index
		fmt.Println(i)
		if i == len(dp) {
			dp = append(dp, num) //num 比dp中的元素都大 应该加到末尾
		} else {
			dp[i] = num //num加到dp【i】处
		}
	}
	//fmt.Println(dp)
	return len(dp)
}

//动态规划 + 二分法
func lengthOfLIS2(nums []int) int {
	var dp []int
	for _, v := range nums {
		//如果v比dp的最后一个元素还大，说明v放到dp后还是符合严格递增的
		if len(dp) == 0 || v > dp[len(dp)-1] {
			dp = append(dp, v)
		} else {
			var left = 0
			var right = len(dp)
			for left < right {
				mid := left + (right-left)/2
				if dp[mid] < v {
					left = mid + 1
				} else {
					right = mid
				}
			}
			//此时left == right 哪个都一样
			dp[left] = v
		}
	}
	return len(dp)
}
func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
	fmt.Println(lengthOfLIS2(nums))
}
