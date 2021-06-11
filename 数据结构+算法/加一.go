package main

import "fmt"

func plusOne1(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println(carry)
		carry, digits[i] = (digits[i]+carry)/10, (digits[i]+carry)%10
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main()  {
	var nums = []int{9,9,9}
	fmt.Println(plusOne1(nums))
}