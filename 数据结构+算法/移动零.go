package main

import "fmt"

func moveZeroes(nums []int) {
	i := 0 //统计前面0的个数
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 { //当前数字是0就不操作
			i++
		} else if i != 0 {
			//否则就把当前数字放到最前面那个0的位置，然后再把当前位置设置成0
			nums[j-i] = nums[j]
			//fmt.Println(j , i) // 2 2,3 2,4 2
			nums[j] = 0
		}
	}
}

func main() {
	var nums = []int{0, 0, 1, 2, 3}
	moveZeroes(nums)
	fmt.Println(nums)
}
