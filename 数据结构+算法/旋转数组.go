package main

import "fmt"

func rotate(nums []int, k int) {
	var length = len(nums)
	var temp = []int{}
	//把原数组值放到一个临时数组中，
	for i := 0; i < length; i++ {
		temp = append(temp, nums[i])
	}
	//然后在把临时数组的值重新放到原数组，并且往右移动k位
	for i := 0; i < length; i++ {
		nums[(i+k)%length] = temp[i]
	}
}

func main() {
    var nums = []int{-1,-100,3,99}
	rotate(nums,2)
    fmt.Println(nums)
}
