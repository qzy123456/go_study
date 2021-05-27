package main

import "fmt"

func finRepeatNumbers(nums []int)int  {

	var same = make(map[int]struct{})
	for _,num := range nums{
		if _,ok:=same[num];ok{
			return num
		}
		same[num] = struct{}{}
	}
	return -1
}

//辅助数组版
func finRepeatNumbers2(nums []int)int  {

	var n  = len(nums)
	var arr  = make([]int,n)
	for i:=0;i<n;i++  {
		arr[nums[i]]++
		if arr[nums[i]] > 1{
			return nums[i]
		}
	}

	return -1
}

func main() {
	nums := []int{1,1,3,2}
	fmt.Println(finRepeatNumbers(nums))
	fmt.Println(finRepeatNumbers2(nums))
}
