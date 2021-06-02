package main

import (
	"fmt"
)

func BinarySort(arr *[]int, leftIndex, rightIndex, findVal int) {
	// 判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		// 说明我们要找的数，应该在leftIndex---middle-1
		BinarySort(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 说明我们要找的数，应该在middle+1---rightIndex
		BinarySort(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("找到了，下标为%v\n", middle)
	}
}

func main() {
	slice := []int{24, 69, 80, 57, 13}
	BinarySort(&slice, 0, len(slice)-1, 80)

}
