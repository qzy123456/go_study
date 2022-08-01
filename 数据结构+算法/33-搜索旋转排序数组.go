package main

import "fmt"

func searchInserts1(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	} else {
		for i, v := range nums {
			if v == target {
				return i
				break
			}
		}
	}
	return -1
}

func searchInserts(arr []int, value int) int {
	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	if len(arr) == 1 {
		if arr[0] != value {
			return -1
		}
	}

	for start <= end {
		//找到了
		if arr[mid] == value {
			return mid
		}
		if arr[0] <= arr[mid] {
			//数字在前半部分
			if arr[0] <= value && arr[mid] > value {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if arr[mid] < value && arr[end] >= value {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		mid = start + (end-start)/2

	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 2
	fmt.Println(searchInserts(nums,target))
}