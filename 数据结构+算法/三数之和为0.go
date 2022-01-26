package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int)[][]int{
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
func threeSum2(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//第2个数
		j := i + 1
		//最后一个
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}

func threeSum3(nums []int) [][]int {
	// 先从小到大排序
	sort.Ints(nums)
	// 接收结果
	var res [][]int
	// 获取数组长度
	length := len(nums)
	// 边界处理，数字不足三个直接返回空
	if len(nums) < 3 {
		return res
	}
	// 开始循环第一个固定值
	for index, _ := range nums {
		// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
		if nums[index] > 0 {
			return res
		}
		// 排除值重复的固定位
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		// 指针初始位置，固定位右边第一个和数组最后一个
		l := index + 1
		r := length - 1
		// 开始移动两个指针
		for l < r {
			// 判断三个数字之和的三种情况
			sum := nums[index] + nums[l] + nums[r]
			switch {
			case sum == 0:
				// 将结果加入二元组
				res = append(res, []int{nums[index], nums[l], nums[r]})
				// 去重，如果l < r且下一个数字一样，则继续挪动
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 同理
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case sum > 0:
				// 如果和大于 0，那就说明 right 的值太大，需要左移
				r--
				// 如果和小于 0，那就说明 left 的值太小，需要右移
			case sum < 0:
				l++
			}
		}
	}
	return res
}

func main() {
	nums := []int{1,-1,0,2,-1,-1}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum2(nums))
	fmt.Println(threeSum3(nums))
}
