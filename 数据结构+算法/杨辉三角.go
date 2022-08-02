package main

import "fmt"

//     1
//    1 1
//   1 2 1
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	yanghui := make([][]int,numRows)
	for i:=0;i<len(yanghui);i++{
		// 此处需要对yanghui[i]进行make分配
		yanghui[i] = make([]int,i+1)
	}
	for i := 0; i < numRows; i++ {   //i是行
		for j := 0; j <= i; j++ {  //j是列
			if i < 2 { //两行以内三角中的数字都是1
				yanghui[i][j] = 1
			} else { //第三行开始，正式计算数值写入数组
				if j == 0 || j == i { //所有行的第一列和最后一列都是1
					yanghui[i][j] = 1
				} else { //当前数组元素是上一行的前一个元素加上上一行的当前列元素
					yanghui[i][j] = yanghui[i-1][j-1] + yanghui[i-1][j]
				}
			}
		}
	}
	return yanghui
}

func main() {
  fmt.Println(generate(5))
}
