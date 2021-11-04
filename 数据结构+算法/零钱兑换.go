package main

import (
	"fmt"
	"sort"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
//可以认为硬币数量是无限的
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	res := make([]int, amount)
	for index := 0; index < len(res); index++ {
		res[index] = -1
	}
	//mt.Println(res) [-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
	for money := 1; money <= amount; money++ {
		for _, coin := range coins {
			lessMoney := money - coin
			if lessMoney < 0 {
				continue
			} else if lessMoney == 0 {
				res[money-1] = 1
			} else {
				coinNums := res[money-coin-1]
				if coinNums == -1 {
					continue
				} else {
					if res[money-1] == -1 {
						res[money-1] = coinNums + 1
					} else {
						if coinNums+1 < res[money-1] {
							res[money-1] = coinNums + 1
						}
					}
				}
			}
		}
	}
	return res[amount-1]
}

func coinChange2(coins []int, amount int) int {
	sort.Ints(coins)
	ret := make([]int, amount+1)
	for k,_:=range(ret){
		ret[k] = amount+1
	}
	fmt.Println(ret) //[12 12 12 12 12 12 12 12 12 12 12 12]
	ret[0] = 0
	for i:=1;i<=amount;i++{
		for j:=0;j<len(coins);j++{
			if i >= coins[j] {
				ret[i] = min(ret[i], ret[i-coins[j]] + 1)
			}
		}
	}
	if ret[amount] == amount+1{
		return -1
	}
	return ret[amount]
}

func min(a, b int) int{
	if a > b{
		return b
	}
	return a
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChange2(coins, amount))
}
