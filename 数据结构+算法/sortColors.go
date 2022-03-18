package main

import "fmt"

//这题由于数字只会出现 0，1，2 这三个数字，所以用游标移动来控制顺序也是可以的。
// 具体做法：0 是排在最前面的，所以只要添加一个 0，就需要放置 1 和 2。
// 1 排在 2 前面，所以添加 1 的时候也需要放置 2 。至于最后的 2，只用移动游标即可。
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
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
	low, mid, high := 0, 0, len(nums)-1
	// 定义三个指针，当中间的指针超过右边的指针时，就返回了
	// 算法核心：
	for mid <= high {
		if nums[mid] == 1 {
			// 当中间指针指向的值是1时，中间指针向后移动一位
			mid++
			continue
		} else if nums[mid] == 0 {
			// 交换中间指针和左边指针的值
			// 当中间指针指向的值是0时，如果此时左边指针和中间指针重合，那么两个指针都向后移动一位
			// 否则只是左边的指针向后移动一位
			nums[low], nums[mid] = nums[mid], nums[low]
			if mid == low {
				mid++
			}
			low++
		} else if nums[mid] == 2 {
			// 交换中间指针和右边指针的值
			// 当中间的指针指向的值是2时，右边的指针向左移动一位
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}

func main() {
	nums :=[]int{1,2,0,1,2}
	sortColors(nums)
	fmt.Println(nums)
}