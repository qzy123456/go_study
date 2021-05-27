package main

import "fmt"

//无序的情况下，先排序

//循环暴力破解
func solution3(nums []int,target int) (int,int)  {
	var length = len(nums)
	for i:=0;i<length-1 ;i++  {
		num := target-nums[i]
		for j:=i+1;j<length ;j++  {
			if nums[j] == num{
				return i,j
			}
		}
	}

	return -1,-1
}
//map打标记
func solution4(nums []int,target int) (int,int)  {
	var arrMap  = make(map[int]int)
	for j:=0;j<len(nums) ;j++  {
		num := target-nums[j]
		if index,ok:=arrMap[num];ok{
            return index,j
		}
		arrMap[nums[j]] = j
	}
	return -1,-1
}
//二分查找
func solution(nums []int,target int) (int,int)  {
	var low  = 0
	var high  = len(nums)-1
	for i:=0;i<len(nums) ;i++  {
		//low设置个默认值
		low = i
		for low < high {
			mid := (low + high)/2
			if nums[mid] == target - nums[i]{
				return i,mid
			}else if nums[mid] > target -nums[i]{
				high = mid -1
			}else {
				low = mid +1
			}
		}
	}
	return -1,-1
}
//双指针
func solution2(nums []int,target int)(int, int)  {
	var low  =  0
	var high  = len(nums)-1
	for low < high {
		switch sum:=nums[high] + nums[low]; {
		case sum == target:
			return low, high
		case sum < target:
			low++
		default:
			high--
		}
	}

	return -1,-1
}

func main() {
	nums := []int{1,2,3,5}
	fmt.Println(solution(nums,3))
	fmt.Println(solution2(nums,3))
	fmt.Println(solution3(nums,3))
	fmt.Println(solution4(nums,3))
}