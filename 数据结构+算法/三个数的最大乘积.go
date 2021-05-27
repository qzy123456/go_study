package main

import (
	"fmt"
	"math"
	"sort"
)

//先排序
//全是正数或全是负数，则nums[length-1]*nums[length-2]*nums[length-3]
//有负数有正数，则要么是nums[length-1]*nums[length-2]*nums[length-3]，要么是nums[0]*nums[1]*nums[length-1]
func find3MaxProduct(nums []int)int  {
	sort.Ints(nums)
	leng := len(nums)
	res := math.Max(float64(nums[0]*nums[1]*nums[leng-1]),float64(nums[leng-1]*nums[leng-2]*nums[leng-3]))
	return int(res)
}

//线性扫描
func find3MaxProduct2(nums []int)int{
	var min1 = math.MaxInt64
	var min2 = math.MaxInt64

	var max1 = math.MinInt64
	var max2 = math.MinInt64
	var max3 = math.MinInt64

	for _, v := range nums{
		if v <min1{
			min2 = min1
			min1 = v
		}else if v<min2{
			min2 = v
		}

		if v>max1{
			max3 = max2
			max2 = max1
			max1 = v
		}else if v>max2{
			max3 = max2
			max2 = v
		}else if v>max3{
			max3 = v
		}
	}
	res := math.Max(float64(min1*min2*max1),
		float64(max1*max2*max3))
	return int(res)
}

func main() {
	nums :=[]int{-11,-12,3,4,5,6}
	fmt.Println(find3MaxProduct(nums))
	fmt.Println(find3MaxProduct2(nums))
}