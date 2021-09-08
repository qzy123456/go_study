package main

import "fmt"
//二分的时候，中点可能落在3种情况中:
//1:nums[l], nums[m], nums[r]单调递增
//2:nums 大小会发生突变。nums[m] 在突变点左侧，即nums[l]< nums[m]
//3:nums 大小会发生突变。nums[m] 在突变点右侧，即nums[l]> nums[m]
//我们考虑要去除左半边的情况。对于 1, 3, 如果满足 nums[m] < target <= nums[r]，那么我们就能去除左半边。
//对于2，我们知道此时 m 的值是最大的。此时 target 想要落在 m 的右侧，他要么比 m 还大(此时 target 在突变左侧)，要么他比 r 还要小。
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	//由于数组是无序的 例如 4,5,6,7,0,1,2 ---- target 是0
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
func search(nums []int, target int) int {
	//最左，最右
	l, r := 0, len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		lv, mv, rv := nums[l], nums[mid], nums[r]
		//目标大于中间的，并且小于末尾的。或者中间的大于最左，中间的大于最右，目标值大于中间或者目标值小于等于最右
		if (mv < target && target <= rv) || (mv >= lv && mv >= rv && (target > mv || target <= rv)) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func main()  {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	nums1 := []int{4,5,6,7,0,1,2}
	target1 := 3

	fmt.Println(search(nums,target))
	fmt.Println(search1(nums1,target1))
}
