package main

import (
	"fmt"
	"time"
)

func main() {
	x := 2
	y := 4
    //定义一个多维数组
	table := make([][]int, x)
	for i  := range table {
		//多维数组的循环赋值
		table[i] = make([]int, y)
	}
	fmt.Println(table) //[[0 0 0 0] [0 0 0 0]]
	h, w := 2, 4
	//数组
	raw := make([]int, h*w)
    //赋值
	for i := range raw {
		raw[i] = i
	}
	// 初始化原始 slice
	fmt.Println(raw, &raw[4])	// [0 1 2 3 4 5 6 7] 0xc420012120

	table1 := make([][]int, h)
	for i := range table1 {
		// 等间距切割原始 slice，创建动态多维数组 table
		// 0: raw[0*4: 0*4 + 4]
		// 1: raw[1*4: 1*4 + 4]
		table1[i] = raw[i*w : i*w + w] //两个slice数组的地址其实是一样的）
	}
	fmt.Println(table1, &table1[1][0])	// [[0 1 2 3] [4 5 6 7]] 0xc420012120 （可以发现 raw的地址跟table1的指针地址是相同的）

	var mSlice1 []string = make([]string,3,5)
	fmt.Printf("mSlice1的长度是%d,容量数%d,内容是%v\n", len(mSlice1), cap(mSlice1), mSlice1)
	var mSlice2 []string = make([]string,3,3)
	mSlice2[0] = "tom"
	mSlice2[1] = "kobe"
	mSlice2[2] = "jack"
	//理论上如果新增数据  大于cap的话  默认cap是2倍增长的
	mSlice2 = append(mSlice2, "xasxas")
	fmt.Printf("mSlice2的长度是%d,容量数%d,内容是%v\n", len(mSlice2), cap(mSlice2), mSlice2)
	//基于切片创建切片
	var oldS []int= make([]int,3,10)
	fmt.Printf("切片oldS的长度是%d,容量是%d,内容是 %v\n", len(oldS), cap(oldS), oldS)
	var newS = oldS[:6]
	fmt.Printf("切片newS的长度是%d,容量是%d,内容是 %v\n", len(newS), cap(newS), newS)
	ch1 := make(chan rune)
	go func() {
		for c := 'a'; c < 'a'+26; c++ {
			ch1 <- c
			time.Sleep(time.Second)
		}
		//关闭channel
		close(ch1)
	}()
	orderLen := 5
	order := make([]uint16, 2 * orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
				//先计算order[orderLen:]，得到【0，0，0，0，0】，再计算[:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}