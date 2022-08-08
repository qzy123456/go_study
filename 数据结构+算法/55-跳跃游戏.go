package main

import "fmt"

//数组的倒数第二个个元素看起，每次往前遍历，如果当前的元素能够到达最后一个位置，那么我们把当前位置开始到最后全部“切断”，
//以当前元素为最后一个元素，重复第一次的过程就好了。而最后，如果我们遍历到数组的第一个元素，那么此题返回true
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
func canJump(nums []int) bool {
	l := len(nums) - 1
	for i := l-1; i >=0 ; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}

func canJump2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if maxJump < i  {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}
func max(a,b int)int{
	if a>b{
		return  a
	}
	return  b
}
func main() {
	nums := []int{2,3,0,1,4}
	fmt.Println(canJump(nums))
	fmt.Println(canJump2(nums))
}