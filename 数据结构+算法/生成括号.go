package main

import "fmt"

//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//例如，n的数量为3，输出如下:
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}
func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

func generateParenthesis2(n int) []string {
	res := make([]string, 0)
	generate("",&res, n, n)
	return res
}

func generate(s string, res *[]string, left int,right int ) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}

	if left > 0 {
		generate(s + "(", res, left - 1, right)
	}
	if right > left {
		generate(s + ")", res, left, right - 1)
	}
}


func main() {
	fmt.Printf("%#v", generateParenthesis(3))
	fmt.Println()
	fmt.Printf("%#v", generateParenthesis2(3))
}
