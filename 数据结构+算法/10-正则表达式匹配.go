package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) == 0 {
		// 边界情况，即s为空，p前两个为 x*
		if len(p) >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		}
		return false
	} else if len(p) == 0 {
		return false
	}

	// p是否为 x* 形式
	var hasStar bool
	if len(p) >= 2 && p[1] == '*' {
		hasStar = true
	}

	// isMatch表示s与p的第一个字符是否匹配
	var isMatched = true
	if s[0] != p[0] && p[0] != '.' {
		isMatched = false
	}

	if hasStar {
		if isMatched {
			// 情况1： 有星且第一个字符匹配，则递归包括2个情况：s去掉第一个字符，p去掉star这两个字符
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		}
		// 情况2：有星且不匹配，则去掉p的前两个字符继续匹配
		return isMatch(s, p[2:])
	} else if !isMatched {
		// 情况3：没星且不匹配，则直接返回不匹配
		return false
	}
	// 情况4：没有星但是匹配，s和p删掉匹配的第一个字符，继续匹配
	return isMatch(s[1:], p[1:])
}

func isMatch2(s string, p string) bool {
	row, col := len(s), len(p)

	// dp 就是核心的状态转移方程，这里注意要+1，是为了空字符串这个边界条件
	// 所以后面的i/j默认都要-1
	dp := make([][]bool, row+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, col+1)
	}
	// 填充dp[0]数组，也就是s为空字符串
	for j := 0; j < col+1; j++ {
		if j == 0 {
			// p为空字符串的情况
			dp[0][0] = true
		} else if p[j-1] == '*' {
			// 如果p[j-1]为*，则可以认为匹配p和p[0:j-2]一样，类似于情况2
			dp[0][j] = dp[0][j-2]
		}
	}

	// 填充整个dp数组，注意i和j在dp中不变，但对应到字符串s/p中都要-1
	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if p[j-1] == '*' {
				if i != 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					// 对应情况1，有星且第一个字符匹配
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					// 对应情况2，有星且不匹配
					dp[i][j] = dp[i][j-2]
				}
			} else if i != 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
				// 对应情况4，没有星但是匹配
				dp[i][j] = dp[i-1][j-1]
			}
			// 其余的对应情况3，没星且不匹配，即默认false
		}
	}

	return dp[row][col]
}

func main()  {
	fmt.Println(isMatch("aa", "b"))
	fmt.Println(isMatch2("aa", "a*"))
}
