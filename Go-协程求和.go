package main

import (
	"fmt"
)

func splitSlice(slice []int, n int) [][]int {
	length := len(slice) / n
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		start := i * length
		end := (i + 1) * length
		if i == n-1 {
			end = len(slice)
		}
		result[i] = slice[start:end]
	}
	return result
}

func chunkSlice(slice []int, n int) [][]int {
	var chunks [][]int
	sliceLen := len(slice)
	chunkSize := (sliceLen + n - 1) / n // round up division

	for i := 0; i < sliceLen; i += chunkSize {
		end := i + chunkSize
		if end > sliceLen {
			end = sliceLen
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}


func sum(numbers []int, resultChan chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	resultChan <- sum // 将结果发送到通道中
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int) // 定义一个通道用于接收结果

	go sum(numbers[:len(numbers)/2], resultChan) // 启动第一个协程计算前一半数的和
	go sum(numbers[len(numbers)/2:], resultChan) // 启动第二个协程计算后一半数的和

	// 在主协程中等待两个协程都完成，并将它们的结果相加
	sum1, sum2 := <-resultChan, <-resultChan
	totalSum := sum1 + sum2

	fmt.Println("Total Sum: ", totalSum)
}
