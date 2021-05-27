package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, k int
	var m = len(nums1)
	var n = len(nums2)
	var nums = make([]int, m+n)

	for i != m && j != n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}
	//一个数组长，一个数组短，那么就会一个没处理完
	for i != m {
		nums[k] = nums1[i]
		i++
		k++
	}
	for j != n {
		nums[k] = nums2[j]
		j++
		k++
	}
	var mid int = int((m + n) / 2)
	//偶数位
	if (m+n)%2 == 0 {
		return float64(nums[mid]+nums[mid-1]) / 2
	}

	return float64(nums[mid])
}

func main() {

	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{5, 6, 7}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
