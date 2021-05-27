package main

import (
	"fmt"
	"sort"
)

//贪心算法  先升序  倒叙查找3个数  看是否满足条件

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i-1] + nums[i-2] + nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{3, 6, 2, 3}
	fmt.Println(largestPerimeter(nums))
}
