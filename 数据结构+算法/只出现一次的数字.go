package main

import "fmt"

func singleNumber(nums []int) int {
   var temp = 0
    for i:=0;i<len(nums);i++{
    	temp = temp^nums[i]
	}
    return temp
}

func main() {
    var nums = []int{1,1,3,4,3}
    fmt.Println(singleNumber(nums))

}