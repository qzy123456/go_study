package main

import "fmt"

//递归
func fibRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

//动态分配
func fibRecursive2(n int) int {
	var arr = []int{0, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}

//动态规划----空间改进
func fibRecursive3(n int) int {
	var arr = []int{0, 1}
	var temp = 0
	for i := 2; i <= n; i++ {
		temp = arr[1]
		arr[1] = arr[0] + arr[1]
		arr[0] = temp
	}
	return arr[1]
}

//双指针
func fibRecursive4(n int) int {
	var low = 0
	if n == 0 {
		return low
	}
	var high = 1
	for i := 2; i <= n; i++ {
		sum := low + high
		low = high
		high = sum
	}
	return high
}

func main() {
	fmt.Println(fibRecursive(3))
	fmt.Println(fibRecursive2(3))
	fmt.Println(fibRecursive3(3))
	fmt.Println(fibRecursive4(3))
}
