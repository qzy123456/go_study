package main

import (
	"fmt"
	"reflect"
	"strconv"
)
//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//按大小顺序列出所有排列情况，并一一标记，
// 当 n = 3 时, 所有排列如下：“123”，“132”，“213”，“231”，“312”，“321”，
// 给定 n 和 k，返回第 k 个排列。
func getPermutation(n int, k int) string {
	factors := make([]int, 10, 10)
	nums := make([]int, 9, 9)
	res := ""
	factors[0] = 1
	nums[0] = 1
	for i := 1; i < 9; i++ {
		factors[i] = factors[i-1] * i
		nums[i] = i + 1
	}
	k--
	for i := n; i >= 1; i-- {
		rs := k / factors[i-1]
		fmt.Println(nums[rs], reflect.TypeOf(string(nums[rs])))
		k %= factors[i-1]
		res += string(nums[rs] + '0')
		nums = append(nums[:rs], nums[rs+1:]...)
	}
	return res
}
//第二种解法
func getPermutation2(n int, k int) string {
	if k == 0 {
		return ""
	}
	used, p, res := make([]bool, n), []int{}, ""
	findPermutation(n, 0, &k, p, &res, &used)
	return res
}

func findPermutation(n, index int, k *int, p []int, res *string, used *[]bool) {
	fmt.Printf("n = %v index = %v k = %v p = %v res = %v user = %v\n", n, index, *k, p, *res, *used)
	if index == n {
		*k--
		if *k == 0 {
			for _, v := range p {
				*res += strconv.Itoa(v + 1)
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, i)
			findPermutation(n, index+1, k, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func main()  {
    fmt.Println(getPermutation(3,3))
    fmt.Println(getPermutation2(3,3))
}
