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
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
		fmt.Println(midIndex1,midIndex2)
		return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}
func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {
	// 获取两个数组的长度
	nums1Length := len(nums1)
	nums2Length := len(nums2)
	// 总长度
	sumLength := nums1Length + nums2Length
	halfLength := 0
	// 确定半长，因为如果总长度是偶数，那么中位数就是两个数字的和除以2，如果总长度是奇数，那么中位数就是一位数字
	if sumLength%2 != 0 {
		halfLength = (sumLength + 1) / 2
	} else {
		halfLength = sumLength / 2
	}
	//fmt.Println(halfLength)
	// 这三个用来存储临时变量，也就是和中位数有关的一位或者两位数字
	tmpNum0, tmpNum1, tmpNum2 := 0, 0, 0
	// 题目要求时间复杂度是O(m+n)，所以使用两个下标来遍历两个数组
	nums1Index, nums2Index := 0, 0
	// 这个循环就是最后找到中位数，循环比总长度的一半多找了后面的一个数字，
	// 这样的目的是当中位数是两个数字的时候，一次性就找出这两个数字，代码看起来会简单一点
	// 最开始我循环的临界是i<halfLength，这样的结果是当中位数是两个数字的时候，后面会有很大一部分重复代码
	for i := 0; i <= halfLength; i++ {
		// 第一个条件判断是当两个数组都还没有越界的时候，所以需要取两个数组中较小的那个数字
		if nums1Length > nums1Index && nums2Length > nums2Index {
			// 取了某个数组中的数字，那个数组的下标就向后加1
			// 使用tmpNum0保存临时值
			if nums1[nums1Index] <= nums2[nums2Index] {
				tmpNum0 = nums1[nums1Index]
				fmt.Println("1:",tmpNum0)
				nums1Index++
			} else {
				tmpNum0 = nums2[nums2Index]
				fmt.Println("2:",tmpNum0)
				nums2Index++
			}
		} else if nums1Length <= nums1Index && nums2Length > nums2Index {
			// 当数组nums1已经发生了越界，那么此时只剩nums2，所以取nums2的数字
			tmpNum0 = nums2[nums2Index]
			fmt.Println("3:",tmpNum0)
			nums2Index++
		} else if nums1Length > nums1Index && nums2Length <= nums2Index {
			// 当数组nums2已经发生了越界，那么此时只剩nums1，所以取nums1的数字\
			fmt.Println("4:",tmpNum0)
			tmpNum0 = nums1[nums1Index]
			nums1Index++
		}
		// 之前取出的数字保存在临时变量tmpNum0中
		// 之前的值一直保存在tmpNum1中，tmpNum1一直在更新，直到更新到最中间的数字，或者中间的两个数字中较小的那个
		// 只有最后一个数字，也就超过半数的那个数字保存tmpNum2中
		if i < halfLength {
			tmpNum1 = tmpNum0
			fmt.Println("5:",tmpNum0)
		} else {
			tmpNum2 = tmpNum0
			fmt.Println("6:",tmpNum0)
		}
	}
	if sumLength%2 != 0 {
		// 此时中位数就只有一个
		return float64(tmpNum1)
	} else {
		// 此时中位数是中间的两个数字的和除以2
		return float64(float64(tmpNum1 + tmpNum2) / 2)
	}
}
func main() {

	nums1 := []int{1, 4, 5}
	nums2 := []int{3, 6, 7}

	//fmt.Println(findMedianSortedArrays(nums1, nums2))
	//fmt.Println(findMedianSortedArrays2(nums1, nums2))
	//fmt.Println(findMedianSortedArrays3(nums1, nums2))
	fmt.Println(findMedianSortedArrays4(nums1, nums2))
}
