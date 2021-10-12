package main

import "fmt"

func canCircle(oil, cost []int) int {
	//carOil汽车当前油量
	//minCarOil汽车油量最低
	//minIndex汽车油量最低出现在哪个加油站
	var carOil, minCarOil, minIndex int
	for i := 0; i < len(oil); i++ {
		carOil = carOil + oil[i] - cost[i]
		if carOil < minCarOil {
			minCarOil = carOil
			//已经到达下个节点
			minIndex = i + 1
		}

	}
	//走了一圈就是把所有油量加起来
	if carOil < 0 {
		return -1
	}
	return minIndex % len(oil)
}
func main() {
	oil  := []int{4, 5, 3, 1, 4, 3}
	cost := []int{5, 4, 3, 4, 2, 1}
	fmt.Println(canCircle(oil, cost))
}