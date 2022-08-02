package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//从后往前遍历，检测到第一个如果前一元素小于当前元素，那就说明下一排列就需要替换前一元素，
//因为该元素之后的序列是一个递减的排列，然后从这个前一元素开始向后遍历，直到遍历到第一个比这个元素小的元素，
//如果都比这个元素大，就取最后的元素，然后与之替换，再然后，反转这个前一元素之后的序列即可。这样，后面的序列又变成递增的了。
func nextPermutation2(nums []int) {
	i := len(nums) - 2 //1
	j := len(nums) - 1  //2
	//从后往前找相对小的数---[小数]
	for ; i >= 0; i-- {
		fmt.Println("nums[i],nums[i+1]",nums[i],nums[i+1])
		if nums[i] < nums[i+1] {
			break
		}
	}
	fmt.Println("i",i) //1
	//如果i >=0 此时[i+1,n]是升序
	if i >= 0 {
		//从[i+1,n]中倒着找第一个大于 [小数] 的数---[大数]
		//为什么是第一个，因为这样较大数就尽可能小了
		for ; j >= i+1; j-- {
			fmt.Println("nums[i],nums[j]",nums[i],nums[j])
			if nums[i] < nums[j] {
				break
			}
		}
		fmt.Println("j",j) //2
		//交换 大数 和 小数
		nums[i], nums[j] = nums[j], nums[i]
		//此时[i+1,n]还是降序
	}
	//如果i=-1说明数组是降序排列，直接改成升序即可
	//将 大数 后面的所有升序
	fmt.Println(nums,i+1,nums[i+1:])
	sort.Ints(nums[i+1:])
}

func main(){
	nums := []int{4,3,2,1}
	nextPermutation2(nums)
	fmt.Println(nums)
}
