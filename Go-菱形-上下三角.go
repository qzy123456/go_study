package main


import "fmt"

func main() {
	// 长
	x := 9
	// 宽
	y := 9
	// 行数
	row := 1
	for row <= y {
		// 计算每行得星星数
		count := 0
		if row <= (y/2 + 1) {
			count = 2 * row - 1
		} else {
			count = 2 * (y - row) + 1
		}
		row++
		text := ""
		// 算出显示星星的范围
		star_min := ((x - count) / 2) + 1
		star_max := ((x - count) / 2) + count
		for index := 1;index <= x;index++ {
			if index >= star_min && index <= star_max {
				text += "*"
			} else {
				text += " "
			}
		}
		fmt.Println(text)
	}
}