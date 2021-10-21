package main

import (
	"fmt"
	"sort"
)

//先按照区间起点进行排序。然后从区间起点小的开始扫描，依次合并每个有重叠的区间。

type slices [][]int
func (s slices) Len() int {
	return len(s)
}
func (s slices) Less(i, j int) bool {
	if s[i][0] <= s[j][0] {
		if s[i][0] == s[j][0] {
			if s[i][1] <= s[j][1] {
				return true
			}
		}
		return true
	}
	return false
}
func (s slices) Swap(i, j int) {
	s[i][0], s[i][1], s[j][0], s[j][1] = s[j][0], s[j][1], s[i][0], s[i][1]
}
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Sort(slices(intervals))

	slow, fast := 0, 1
	for fast != len(intervals) {
		// 如果不需要合并，那么slow和fast都向后移一步，并且将当前的slow存入res
		if intervals[slow][1] < intervals[fast][0] {
			res = append(res, intervals[slow])
			slow = fast
			fast++
		} else {
			// 需要合并，那么就合并，并且slow不动，fast向后移动一步
			var bigger int
			if intervals[slow][1] >= intervals[fast][1] {
				bigger = intervals[slow][1]
			} else {
				bigger = intervals[fast][1]
			}
			intervals[slow][1] = bigger
			fast++
		}
		if fast == len(intervals) {
			res = append(res, intervals[slow])
		}
	}
	return res
}

func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	//排完序，取出数组第一个元素
	res = append(res, intervals[0]) //[1,6]
    //从第二个开始遍历
	for i := 1; i < len(intervals); i++ {
		l := len(res)
		//前面数组的第2个元素，小于后面数组的第一个元素，证明没交叉
		if res[l-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			//前面数组的第2个元素，小于后面数组的第1个元素，但是前面数组的第2个元素，小于后面数组的第二个元素，有交叉
			if res[l-1][1] < intervals[i][1] {
				res[l-1][1] = intervals[i][1]
			}
		}
	}
	return res
}

func main() {
	nums := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(nums))
	fmt.Println(merge2(nums))
}