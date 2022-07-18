package main

import "fmt"

// 堆排序
func HeapSort(data []int) {
	// 构建堆
	length := len(data)
	for i := (length - 2) / 2; i >= 0; i-- {
		heapify(data, length, i)
	}

	// 排序
	for length > 0 {
		length--
		data[length], data[0] = data[0], data[length]
		heapify(data, length, 0)
	}
}

func heapify(data []int, size, i int) {
	for {
		c1 := 2*i+1
		c2 := 2*i+2
		max := i
		if c1 < size && data[c1] > data[max] {
			max = c1
		}
		if c2 < size && data[c2] > data[max] {
			max = c2
		}
		if i == max {
			break
		}
		data[i], data[max] = data[max], data[i]
		i = max
	}
}

func main() {
	ar := []int{3,1,5,4,9,0}
	HeapSort(ar)
	fmt.Println(ar)
}
