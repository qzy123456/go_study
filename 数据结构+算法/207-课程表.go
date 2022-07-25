package main

import "fmt"

//现在你总共有 n 门课需要选，记为 0 到 n-1。在选修某些课程之前需要一些先修课程。
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]。
// 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
//解题思路 #
//给出 n 个任务，每两个任务之间有相互依赖关系，比如 A 任务一定要在 B 任务之前完成才行。问是否可以完成所有任务。
//这一题就是标准的 AOV 网的拓扑排序问题。拓扑排序问题的解决办法是主要是循环执行以下两步，直到不存在入度为0的顶点为止。
//选择一个入度为0的顶点并输出之；
//从网中删除此顶点及所有出边。
//循环结束后，若输出的顶点数小于网中的顶点数，则输出“有回路”信息，即无法完成所有任务；否则输出的顶点序列就是一种拓扑序列，即可以完成所有任务。
// AOV 网的拓扑排序
func canFinish(n int, pre [][]int) bool {
	in := make([]int, n)
	frees := make([][]int, n)
	next := make([]int, 0, n)
	for _, v := range pre {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	return len(next) == n
}
//深度优先
func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result []int
		valid = true
		dfs func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
//广度优先
func canFinish3(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		indeg = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
func main() {
	numCourses := 2
	prerequisites := [][]int{{1,0}}
	fmt.Println(canFinish(numCourses,prerequisites))
	fmt.Println(canFinish2(numCourses,prerequisites))
	fmt.Println(canFinish3(numCourses,prerequisites))
}

