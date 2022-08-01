package main

import "fmt"

//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
//示例 1：
//
//输入：nums = [1,1,1], k = 2
//输出：2
//示例 2：
//
//输入：nums = [1,2,3], k = 3
//输出：2
//此题不能使用滑动窗口来解。因为 nums[i] 可能为负数。
//前缀和的思路可以解答此题，但是时间复杂度有点高了，O(n^2)。考虑优化时间复杂度。
//题目要求找到连续区间和为 k 的子区间总数，即区间 [i,j] 内的和为 K ⇒ prefixSum[j] - prefixSum[i-1] == k。
//所以 prefixSum[j] == k - prefixSum[i-1] 。这样转换以后，题目就转换成类似 A + B = K 的问题了。
//LeetCode 第一题的优化思路拿来用。用 map 存储累加过的结果。如此优化以后，时间复杂度 O(n)

func subarraySum(nums []int, k int) int {
	var sumMap = make(map[int]int)
	var sum int
	var count int

	sumMap[0] = 1
	for _, num := range nums{
		sum += num
		diff := sum - k
		if times, ok := sumMap[diff];ok{
			count += times
		}
		sumMap[sum]++ // 本语句要放到if计数后面 否则会出现k=0时 每次把当前遍历的数组统计进去 比如[-1] k = 0 答案应该是0
		// 如果先sumMap[sum]++ 统计时diff = -1 把nums[0:1]也统计进去了 实际上这个答案是不对的
	}

	return count
}


//枚举
func subarraySum1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{1,1,1}
	k := 2
	fmt.Println(subarraySum(nums,k))
	fmt.Println(subarraySum1(nums,k))
}

