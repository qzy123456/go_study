package main

import "fmt"

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。
//说明：你不能倾斜容器。
//使用双指针，一个指向最左，一个指向最右，维护一个当前的最大值。
//
//思路重点：将两个指针中较小的那一个往中间移动。
//
//背后原理：当前已经有了一个最大值了，可以看做是当前信息条件下的最优。
//而能够优化这个问题的办法就是将限制当前值的条件进行改变，尝试寻找更优的解法。即，将较小的指针往中间移动。
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
   arr :=[]int{1,8,6,2,5,4,8,3,7}
   fmt.Println(maxArea(arr))
}
