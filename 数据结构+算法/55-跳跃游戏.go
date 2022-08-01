package main

import "fmt"

//给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//解题思路 #
//要求找到最少跳跃次数，顺理成章的会想到用贪心算法解题。扫描步数数组，维护当前能够到达最大下标的位置，记为能到达的最远边界，
//如果扫描过程中到达了最远边界，更新边界并将跳跃次数 + 1。
//扫描数组的时候，其实不需要扫描最后一个元素，因为在跳到最后一个元素之前，能到达的最远边界一定大于等于最后一个元素的位置，
//不然就跳不到最后一个元素，到达不了终点了；如果遍历到最后一个元素，说明边界正好为最后一个位置，最终跳跃次数直接 + 1 即可，
//也不需要访问最后一个元素。
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
//上面题型的变种---第45题
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
	fmt.Println(jump(nums))
	fmt.Println(canJump(nums))
	fmt.Println(canJump2(nums))
}