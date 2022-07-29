package main

import "fmt"

//给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//示例 2:
//输入: prices = [1]
//输出: 0
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n  := len(prices)
	f0 := -prices[0]
	f1 := 0
	f2 := 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2 - prices[i])
		newf1 := f0 + prices[i]
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2] - prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	//输入: prices = [1,2,3,0,2]
	//输出: 3
	//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
	prices := []int{1,2,3,0,2}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
}