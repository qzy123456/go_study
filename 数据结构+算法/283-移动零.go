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
			nums[j] = 0
		}
	}
}
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	var nums = []int{0, 0, 1, 2, 3}
	moveZeroes2(nums)
	fmt.Println(nums)
}
