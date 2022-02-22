package main

import "fmt"

//移除数组中的一个元素
func removeElement(arr []int, ele int) []int {
	//如果是空切片，那就返回0
	if len(arr) == 0 {
		return []int{}
	}
	//用一个索引
	//循环去比较
	//当一样的时候就删除对应下标的值
	//当不一样的时候，索引加1
	index := 0
	for index < len(arr) {
		if arr[index] == ele {
			arr = append(arr[:index], arr[index+1:]...)
		} else {
			index++
		}
	}

	return arr
}

func main() {
	arr := []int{1, 2, 2, 3, 3}
	fmt.Println(removeElement(arr, 2))

}
