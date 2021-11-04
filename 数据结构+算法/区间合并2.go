package main

import "fmt"

//可以分 3 段处理，先添加原来的区间，即在给的 newInterval 之前的区间。然后添加 newInterval ，
//注意这里可能需要合并多个区间。最后把原来剩下的部分添加到最终结果中即可
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) < 1 {
		return append(intervals, newInterval)
	}
	var insertIntls [][]int
	var res [][]int

	//先把要插入的数组，循环处理进去。最后再处理最新的
	for i, v := range intervals {
		if intervals[i][0] > newInterval[0] {
			insertIntls = append(insertIntls, newInterval)
			insertIntls = append(insertIntls, v)
		} else {
			insertIntls = append(insertIntls, v)
			if i == len(intervals)-1 {
				insertIntls = append(insertIntls, newInterval)
			}
		}
	}
	insertIntls = append(insertIntls, intervals[0:]...) ////[[1 3] [2 5] [6 9] [1 3] [6 9]] su
	res = append(res, insertIntls[0])
	for i := 1; i < len(insertIntls); i++ {
		l := len(res)
		if res[l-1][1] < insertIntls[i][0] {
			res = append(res, insertIntls[i])
		} else {
			if res[l-1][1] < insertIntls[i][1] {
				res[l-1][1] = insertIntls[i][1]
			}
		}
	}
	return res
}

func main() {
	nums := [][]int{{1,3},{6,9}}
	nums2 := []int{2,5}
	fmt.Println(insert(nums,nums2))
}
