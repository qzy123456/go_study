package main

import "fmt"

//希尔排序其实是对插入排序的一种优化，回想一下，插入排序的流程是：
// 将数据分为了已排序区间和未排序区间，依次遍历未排序区间的值，将其插入到已排序区间合适的位置。
//
//插入排序的一个最大的缺点是：每次只能移动一位，这样在一些极端的情况下会非常低效；
// 例如数据 2 3 5 7 9 0，如果将 0 移动至元素头部，需要遍历整个数组。
//
//希尔排序的优化点就在于此，它的核心思想是将数据中的元素分为了多个组，每一组分别进行插入排序。
//
//举一个简单的例子：有数据 35 33 42 10 14 19 27 44，首先将数据以其长度的 1/2 （也就是 4）为步长，
// 分为了四个组，分别是 {35,14}、{33,19}、{42,27}、{10,44}。
func ShellSort(data []int) {
	length := len(data)
	step := length / 2
	for step >= 1 {
		for i := 0; i < length-step; i++ {
			j := i+step
			k := data[j]
			//fmt.Println(j,step)
			//3 3
			//4 3
			//5 3
			//1 1
			//2 1
			//3 1
			//4 1
			//5 1
			for ; j > step-1 && data[j-step] > k; j -= step {
				data[j] = data[j-step]
			}
			data[j] = k
		}
		step /= 2
	}
}

// 增量序列折半的希尔排序
func ShellSort2(list []int) {
	// 数组长度
	n := len(list)
	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			// 开始插入排序，每一轮的步长为 step
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				// 开始插入排序，每一轮的步长为 step
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

func main() {
	ar := []int{3,1,5,4,9,0}
	ShellSort(ar)
	ShellSort2(ar)
	fmt.Println(ar)
}
