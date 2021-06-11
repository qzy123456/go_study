package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var temp = make([]int, 0, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		//数组1的值大于数组2的值 ，那么数组2要往后移动，同理小的时候，数组1移动
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			temp = append(temp, nums1[i])
			i++
			j++
		}

	}
	return temp
}

func main() {
	var nums1 = []int{4, 9, 5}
	var nums2 = []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(nums1, nums2))
}
