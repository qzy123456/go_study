package main

import "fmt"

func cattle(num int) int {
	if num <= 3 {
		return 1
	} else {
		return cattle(num-1) + cattle(num-3)
	}
}

//非递归
func cattle2(num int) int {
	if num <= 3 {
		return 1
	}
	resArr := make(map[int]int)
	resArr[1] = 1
	resArr[2] = 1
	resArr[3] = 1
	for i := 4; i <= num; i++ {
		resArr[i] = resArr[i-1] + resArr[i-3]
	}
	return resArr[num]
}

func main() {

	fmt.Println(cattle(6))
	fmt.Println(cattle2(6))
}
