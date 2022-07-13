package main

import "fmt"

//从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，
// 高度为数组元素的值。如果下雨了，问这样一个容器能装多少单位的水？
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		// 左边是短板，就要移动左边
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}

func max(a,b int) int {
	if a<b{
		return b
	}
	return a
}

func trap3(height []int) int {
	n := len(height)
	// 小于3格则不能存水
	if n < 3 {
		return 0
	}
	//最左边的值，最右边的值
	leftMax,rightMax := height[0],height[n-1]
	ans := 0
	for l,r:=1,n-2;l<=r;{
		// 短板一端决定了能装多少水
		if leftMax < rightMax { // 左边是短板，就要移动左边
			ans += max(0,leftMax - height[l]) // 把有效雨水量加入结果
			leftMax = max(leftMax,height[l]) // 每次更新左边最大值
			l++
		}else{ // 右边是短板，移动右边，同理
			ans += max(0,rightMax- height[r])
			rightMax = max(rightMax,height[r])
			r--
		}
	}
	return ans
}

func main() {
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))
	fmt.Println(trap3(nums))
}