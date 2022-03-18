package main

import "fmt"

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	//  n - (k - c.size()) + 1
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

func combine2(n int, k int) [][]int {
	if n <= 0 || n < k {
		return [][]int{}
	}
	res := [][]int{}

	graph := make([]int, n)
	for i := range graph {
		graph[i] = i + 1
	}
	// [1,2,3,4]
	dfs(graph, k, 0, []int{}, &res)

	return res
}

func dfs(graph []int, target int, index int, path []int, paths *[][]int) {
	//满了
	if len(path) == target {
		tmp := make([]int, target)
		copy(tmp, path)
		*paths = append(*paths, tmp)
		return
	}

	for i, v := range graph {
		if i < index {
			continue
		}
		//fmt.Println(path, append(path, v),paths)
		dfs(graph, target, index+1, append(path, v), paths)
		index++
	}
}


func main() {
	//fmt.Println(combine(4,2))
	fmt.Println(combine2(4,2))
}
