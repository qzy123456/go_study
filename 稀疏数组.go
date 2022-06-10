package main

import "fmt"

// 数据结构
// 使用稀疏数组来保存期盼
// 存盘
// 恢复

type Node struct {
	row int // 行
	col int // 列
	val int // 值
}

func main() {

	// 1. 创建一个原属数组 1代表黑棋 2代表蓝期
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[1][1] = 1
	chessMap[1][3] = 1
	chessMap[1][4] = 1
	chessMap[1][5] = 1
	chessMap[2][3] = 2

	// out fmt
	//for _, v := range chessMap {
	//  for _, v2 := range v {
	//      fmt.Printf("%d\t", v2)
	//  }
	//  fmt.Println()
	//}

	// 3.转换成稀疏数组, go用结构体保存比较好,
	// 思路--> 1.遍历chessMap,如果我们发现,有一个元素的值不等于0,我们就创建一个node结构体
	//        2.将其放置到对应的切片中

	var sparseArr []Node

	// 标准的一个稀疏数组，应该含有 行列的总值，和默认值 （有多少行，有多少列，他的默认值是什么）
	valNode := Node{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			// 如果这个值不为0，需要记录
			if v2 != 0 {
				// 创建一个节点(值节点）
				valNode := Node{
					row: i,
					col: j,
					val: v2,
				}
				// 将这个值放到稀疏数组
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("输出稀疏数组,当前的稀疏数组是:")
	for i, val := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, val.row, val.col, val.val)
	}

	// 如何恢复呢？将稀疏数组恢复为二纬数组.
	// 先创建一个原始数组,
	var chessMap2 [11][11]int

	for i, v := range sparseArr {
		if i == 0 {
			continue
		}
		chessMap2[v.row][v.col] = v.val
	}

	// 验证是否恢复👌
	fmt.Println("恢复过后的原始数据为:")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
