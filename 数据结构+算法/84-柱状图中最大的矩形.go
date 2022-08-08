package main

import "fmt"

//用单调栈依次保存直方图的高度下标，一旦出现高度比栈顶元素小的情况就取出栈顶元素，
// 单独计算一下这个栈顶元素的矩形的高度。然后停在这里(外层循环中的 i–，再 ++，就相当于停在这里了)，
// 继续取出当前最大栈顶的前一个元素，即连续弹出 2 个最大的，以稍小的一个作为矩形的边，宽就是 2 计算面积…………
// 如果停在这里的下标代表的高度一直比栈里面的元素小，就一直弹出，取出最后一个比当前下标大的高度作为矩形的边。
// 宽就是最后一个比当前下标大的高度和当前下标 i 的差值。计算出面积以后不断的更新 maxArea 即可。
func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			// pop stack
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
		}
		// push stack
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i] - left[i] - 1) * heights[i])
	}
	return ans
}

func largestRectangleArea3(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var stack []int
	Max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[idx]
			peek := 0
			w := i
			if len(stack) != 0 {
				peek = stack[len(stack)-1]
				w = i - peek - 1
			}
			Max = max(Max, w*h)
		}
		stack = append(stack, i)
	}
	return Max
}
//正确的解题方法是构造一个递增的栈,每当有新的数据进来时.进行比较,如果比栈顶小的话,就出栈算面积.
//值得注意的地方有两个:
//1, 就是低高低的情况,所以矩形的长度还要向前计算.所以矩形宽度还要看前一个的索引.如果已经是第一个就要全部计算在内.
//2,就是连续相同高度的情况,如果出现这种情况解决方法有2,一个是让其进栈,一个是永远只取最后一个值进行计算.
//为了便于计算我在这个序列的尾巴增加了一个0高度.
func largestRectangleArea4(heights []int) int {
	result := 0
	tmpResult := 0

	heights = append(heights, 0)
	heightIndex := make([]int, 0) //递增栈
	tmpValue := 0
	tmpIndex := 0

	for index, value := range heights {
		for {
			if len(heightIndex) == 0 || value >= heights[heightIndex[len(heightIndex) - 1]] {
				heightIndex = append(heightIndex, index)
				break
			} else {
				if len(heightIndex) == 1 {
					tmpIndex = -1
				} else {
					tmpIndex = heightIndex[len(heightIndex) - 2]
				}

				tmpValue = heights[heightIndex[len(heightIndex) - 1]]
				tmpResult = (index - tmpIndex - 1) * tmpValue
				if result < tmpResult {
					result = tmpResult
				}

				heightIndex = heightIndex[:len(heightIndex) - 1]
			}
		}
	}
	return result
}

func largestRectangleArea5(heights []int) int {
	//单调栈（单调递增）
	stack := make([]int, 0)
	stack = append(stack, -1) //stack的哨兵，方便确定左边界
	heights = append(heights,0) //添加一个哨兵，减少代码量
	ln := len(heights)
	res := 0 //结果

	for i:=0; i<ln; i++ {
		//因为我们无法访问heights[-1]，所以限制len(stack) > 1
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//栈顶元素，也就是当前要求的矩形柱子的下标
			top := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//左边界（栈顶元素的后一个元素）
			l := stack[len(stack)-1]
			//矩形面积：(右边界-左边界-1) * 高度
			//右边界就是i
			//高度就是以栈顶元素为下标的柱子的高度
			//左边界就是栈顶元素的下一位元素（因为我们添加了哨兵-1，所以这公式依旧成立）
			res = max(res, (i-l-1)*heights[top])
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	heights := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(heights))
	fmt.Println(largestRectangleArea2(heights))
	fmt.Println(largestRectangleArea3(heights))
	fmt.Println(largestRectangleArea4(heights))
	fmt.Println(largestRectangleArea5(heights))
}

