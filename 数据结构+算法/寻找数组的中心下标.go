package main

import "fmt"

func piotIndex(nums []int)int{
	sum :=0 //中心下标右侧和
	total := 0 //中心下标左侧和
    for _,v := range nums{
    	sum += v
	}
	for i:=0;i<len(nums) ;i++  {
		total +=nums[i]
		if total == sum{
			return i
		}
		//不想等sum减去nums【i】
		sum -= nums[i]
	}
	return -1
}

func main() {
    nums := []int{1,7,3,6,5,6}
	fmt.Println(piotIndex(nums))
}