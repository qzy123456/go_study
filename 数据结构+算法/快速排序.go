package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		left, right := start, end
		mid := arr[(start+end)/2]
		for left <= right {
			for arr[left] < mid {
				left++
			}
			for arr[right] > mid {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		if start < right {
			quickSort(arr, start, right)
		}
		if end > left {
			quickSort(arr, left, end)
		}
	}
}
func quick(arr []int,start,end int)  {
	for start<end{
		left := start
		right := end
		mid := arr[(left+end)/2]
		for left <= end{
			if arr[left] < mid{
				left++
			}
			if arr[end] > mid{
				end--
			}
			if left <= end{
				arr[left],arr[end] = arr[end],arr[left]
				left++
				end--
			}
		}
		if start < right{
			quick(arr,start,right)
		}
		if end > left{
			quick(arr,left,end)
		}

	}
}
func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
