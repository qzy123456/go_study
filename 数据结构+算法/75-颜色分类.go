package main

import "fmt"

//这题由于数字只会出现 0，1，2 这三个数字，所以用游标移动来控制顺序也是可以的。
// 具体做法：0 是排在最前面的，所以只要添加一个 0，就需要放置 1 和 2。
// 1 排在 2 前面，所以添加 1 的时候也需要放置 2 。至于最后的 2，只用移动游标即可。
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		//默认第一个全部设置成2
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
func sortColors2(nums []int)  {
	l := 0
	r := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 && i < r {
			nums[r], nums[i] = nums[i], nums[r]
			r--
			i--
		}
	}
}

func sortColors3(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func main() {
	nums :=[]int{1,2,0,1,2}
	sortColors(nums)
	fmt.Println(nums)
}