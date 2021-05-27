package main

import "fmt"

//一个有序数组nums，原地删除重复出现的元素，使每个元素只能出现一次，返回删除后数组的新长度。
//不能使用额外的数组空间，必须在原地修改输入的数组，并在使用O(1)额外空间的条件下完成（临时变量、指针等）

//双指针法
func removeDuplicate(nums []int)int  {
	var i  = 0
	if len(nums) == 0{
		return 0
	}
	//j指到数组末尾的时候，返回i的位置，即0。。。i共i+1个数
	for j:=1;j<len(nums) ;j++  {
		// i和j处的数相同时，只管把j往后移，此时j走得快
		// i和j处的数不同时，i往后移，同时把j处的数字替换过来
		if nums[j] != nums[i]{
			i++
			nums[i] = nums[j]
		}
	}
	return  i+1
}

func main() {
	nums := []int{0,1,2,2,3,3,3,4}
	fmt.Println(removeDuplicate(nums))
}