package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var (
		ans = make([]int, len(nums))
		left = make([]int, len(nums))
		right = make([]int, len(nums))
	)
	left[0] = 1
	right[len(nums) - 1] = 1
	//fmt.Println(left)  // 1 0 0 0
	//fmt.Println(right) // 0 0 0 1
	for i := 1; i < len(nums); i++ {
		//fmt.Println(left[i - 1] , nums[i -  1]) // 1 1, 1 2, 2 3
		left[i] = left[i - 1] * nums[i - 1]
	}
	//fmt.Println(left) // 1 1 2 6
	for i := len(nums) - 2; i >= 0; i-- {
		//fmt.Println(right[i + 1],nums[i + 1]) // 1 4,4 3, 12 2
		right[i] = right[i + 1] * nums[i + 1]
	}
	//fmt.Println(right) //24 12 4 1
	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
func productExceptSelf2(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	right := 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i - 1] * nums[i - 1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right *= nums[i + 1]
		ans[i] *= right
	}
	return ans
}
//动态规划
func productExceptSelf3(nums []int) []int {
	dp := make([]int,len(nums))
	l := 1
	r := 1
	for i:=0;i<len(nums);i++  {
		dp[i] = l
		l *= nums[i]
	}
	fmt.Println(dp)
	for j:=len(nums)-1;j>=0 ;j--  {
		dp[j] = dp[j] * r
		r *= nums[j]
	}
	return dp;
}
func main() {
      arr := []int{1,2,3,4}
      //fmt.Println(productExceptSelf(arr))
	  fmt.Println(productExceptSelf2(arr))
      fmt.Println(productExceptSelf3(arr))
}