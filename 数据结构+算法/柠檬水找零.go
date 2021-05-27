package main

//在柠檬水摊上，每一杯柠檬水的售价为5美元。
//顾客排队购买，一次购买一杯。
//
//每一个顾客只买一杯柠檬水，然后向你付5美元、10美元或20美元。
//必须给每个顾客正确找零。
//注意：一开始你手头并没有任何零钱。
//如果你能给每位顾客正确找零，则返回true，否则返回false。
//
//贪心算法
//顾客给你5块，不用找零
//顾客给你10块，找5块
//顾客给你20块，找10+5块或者5+5+5块（优先找10+5块，剩下5+5+5块，因为5块多了更利于找零）

import (
	"fmt"
)

func change(nums []int) bool {
	var five = 0
	var ten = 0
	for _, bill := range nums {
		switch bill {
		case 5:
			five++ // 刚刚好
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	var nums1 = []int{5, 5, 20} //false
	var nums = []int{5, 5, 10}  //true
	fmt.Println(change(nums))
	fmt.Println(change(nums1))
}
