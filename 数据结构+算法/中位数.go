package main

import "fmt"
// 合并成2个数组
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
//单个数组
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小，防止下标越界，要用短的在前
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	//数组1的中位数
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	//奇数
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	//数组2的中位数
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a,b int)int  {
	if a> b {
		return a
	}
	return  b
}
func min(a,b int)int  {
	if a> b {
		return b
	}
	return  a
}
func main() {

	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{5, 6, 7}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArrays2(nums1, nums2))
}
