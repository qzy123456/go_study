package main

import "fmt"

func missingNumber(nums []int) int {
	var num = 0
	for i:=0;i<=len(nums);i++{
		num += i
	}
	for i:=0;i<len(nums);i++{
		num -= nums[i]
	}
	return num
}

func missingNumber2(nums []int) int {
	var num = 0

	for i:=0;i<len(nums);i++{
		num = num ^ nums[i] ^(i+1)
	}
	return num
}

func missingNumber3(nums []int) int {
	var length = len(nums)
	var  sum = (0 + length) * (length + 1) / 2
	for i := 0; i < length; i++ {
		sum -= nums[i]
	}
	return sum
}
func main() {
	var nums  = []int{0,1,3}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber2(nums))
	fmt.Println(missingNumber3(nums))
}
