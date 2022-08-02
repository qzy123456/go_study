package main

import (
	"fmt"
)
//冒泡
func BubbleSort(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			//从小到大
			if arr[i] > arr[j] {
				//交换数据
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//插入排序
func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		//要比较的值，也就是从第二个跟第一个比开始
		val := arr[i]
		//跟前一个开始比
		index := i - 1

		for index >= 0 && arr[index] > val {
			//fmt.Println(index,arr[index],val)
			arr[index+1] = arr[index]
			index--
		}
		if index + 1 == i {
			continue
		}
		arr[index+1] = val
	}
}
//选择排序
func SelectSort(values *[5]int) {
	length := len(values)

	for i := 0; i < length; i++ {
		min := i // 初始的最小值位置从0开始，依次向右

		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		//fmt.Printf("i:%d min:%d\n", i, min)

		// 把每次找出来的最小值与之前的最小值做交换
		if min != i {
			values[i], values[min] = values[min], values[i]
		}
		//fmt.Println(values)
	}
}
//快速排序，3路快排
func quickSort(a []int, left int, right int) {
	if left >= right {  //一定是left >= right
		return
	}
	//临界值
	temp := a[left]
	//开始的点
	start := left
	//结束的点
	stop := right
	//如果右边一直比左边还大
	for right != left {
		for right > left && a[right] >= temp  {
			right --
		}
		for left < right && a[left] <= temp  {
			left ++
		}
		if right > left {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[right], a[start] = temp, a[right]
	quickSort(a, start, left)
	quickSort(a, right+1, stop)
}
//快速排序，普通
func quickSort2(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}
func main() {
	arr := [5]int{1,43,5,94,90}
	BubbleSort(&arr)
	fmt.Println(arr)

	InsertSort(&arr)
	fmt.Println(arr)

	SelectSort(&arr)

	fmt.Println(arr)

	a := []int{12,3,111,23,65,45}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}