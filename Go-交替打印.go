package main

import (
	"fmt"
	"sync"
)

//使用两个 goroutine 交替打印序列，
// 一个 goroutine 打印数字，
// 另外一个 goroutine 打印字母，
// 最终效果如下：
//
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func main() {
	//
	wg := sync.WaitGroup{}
	wg.Add(2)
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []string{"A", "B", "C", "D", "E", "F"}
	c := make(chan int, 1)
	d := make(chan string, 1)
	go func() {

		i := 0
		for i < 6 {
			fmt.Print(arr1[i])
			fmt.Print(arr1[i+1])
			c <- 1
			i += 2
			<-d
		}
		wg.Done()
	}()

	go func() {
		i := 0
		for i < 6 {
			<-c
			fmt.Print(arr2[i])
			fmt.Print(arr2[i+1])
			i += 2
			d <- ""
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println()
}
