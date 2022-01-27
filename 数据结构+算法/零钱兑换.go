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
	for k :=range ret{
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

func CoinChange3(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	for i := 1; i <= amount; i++{
		dp[i] = i       // dp[i] 最大值为i
		found := false  // 表示dp[i] 是否有效
		for _, coin := range coins{
			if coin <= i && dp[i-coin] >= 0{
				dp[i] = min(dp[i], dp[i-coin] + 1)
				found = true
			}
		}

		if !found {
			dp[i] = -1
		}
	}

	return dp[amount]
}

//可以类比为上台阶问题，有3种面额的硬币{1,2,5}凑成11所用的最少的硬币个数，
// 类比为每次只能上1个，2个或5个台阶，走到第11个台阶的走法
//动态规划，设需要凑成的金额为amount，dp[i]表示凑成i金额所用最少的硬币数，
// 则i的取值范围为1->amount
//状态转义方程，dp[i] = min(dp[i-coins[j]]) + 1,
// j的取值范围为0->len(conins),即需要遍历coins中的每一种硬币，找到最dp[i-coins[j]]小的值,
// 然后再加1即为dp[i]
func coinChange4(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var dp = make(map[int]int)
	//初始化dp数组，组成每一种硬币面额所用的最少硬币数都为1
	for i:=0; i<len(coins); i++ {
		dp[coins[i]] = 1
	}
	for i:=1; i<=amount; i++ {
		//拟定一个最值，若最值没有变化则说明不能凑成该面额
		min := 99999
		for j:=0; j<len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] < min{
				min = dp[i-coins[j]]
			}
		}
		dp[i] = min + 1
	}
	if dp[amount] > 99999 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChange2(coins, amount))
}
