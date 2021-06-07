package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func main()  {
	  nums := []int{1,2,3,4}
	 obj := Constructor(nums)
	 param_1 := obj.Reset()
	 param_2 := obj.Shuffle()
	 fmt.Println(param_1)
	 fmt.Println(param_2)
}