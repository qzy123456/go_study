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

		//start  end
		//left   right
		if start < right {
			quickSort(arr, start, right)
		}
		if left < end {
			quickSort(arr, left, end)
		}
	}
}

func QuickSort(arr []int) []int  {
	if len(arr)<2{
		return arr
	}
	mid := arr[0]
	var left,right []int
	for i :=1;i<len(arr);i++{
		if arr[i]>mid{
			right = append(right,arr[i])
		}else{
			left = append(left,arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	left = append(left,mid)
	return append(left,right...)
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
