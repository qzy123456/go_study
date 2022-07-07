package main

import (
"fmt"
)

func SetWay(myMap *[8][7]int, i int, j int) bool {
	//分析什么情况下就找到通路
	if myMap[6][5] == 2 {
		return true
	} else {
		//如果是可以探测的
		if myMap[i][j] == 0 {
			//假设是通的
			myMap[i][j] = 2
			//依据下右上左进行探测
			if SetWay(myMap, i+1, j) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {
				//死路
				myMap[i][j] = 3
				return false
			}
		} else {
			//否则不能探测
			return false
		}
	}
}

func main() {
	//0：代表没有走过的路
	//1：代表墙
	//2：代表是一个通路
	//3：代表走过，但是不通
	var myMap [8][7]int
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[5][2] = 1
	myMap[5][3] = 1
	myMap[5][4] = 1
	myMap[5][5] = 1
	for i := 0; i < 8; i++ {
		fmt.Println(myMap[i])
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕后的地图：")
	for i := 0; i < 8; i++ {
		fmt.Println(myMap[i])
	}
}
