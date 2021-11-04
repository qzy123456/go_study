package main

import "fmt"

//  假设是3
//  1--2--3
//  8--9--4
//  7--6--5
//给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
func generateMatrix(n int) [][]int {
	mtx := make([][]int, n)
	for r := 0; r < n; r++ {
		mtx[r] = make([]int, n)
	}
	r, c, dr, dc := 0, 0, 0, 1

	//9个图
	for i := 0; i < n*n; i++ {
		mtx[r][c] = i + 1

		if mtx[fix(r+dr, n)%n][fix(c+dc, n)%n] > 0 {
			dr, dc = dc, -dr
		}
		r, c = r+dr, c+dc
	}

	return mtx
}

func fix(x, n int) int {
	if x < 0 {
		return x + n
	}
	return x
}

func generateMatrix2(n int) [][]int {
	// write code here
	retSlice := make([][]int, n)
	for i := 0; i < n; i++ {
		retSlice[i] = make([]int, n)
	}
	//往结构体里写入数字，每次填一个回型
	num := 1
	//i代表回型的长度
	for i := n; i > 0; i = i - 2 {
		//写回型第一行
		for j := (n - i) / 2; j < n-(n-i)/2; j++ {
			retSlice[(n-i)/2][j] = num
			num++
		}
		//写回型最后一列
		for j := (n-i)/2 + 1; j < n-(n-i)/2; j++ {
			retSlice[j][n-(n-i)/2-1] = num
			num++
		}
		//写回型最后一行
		for j := n - (n-i)/2 - 2; j >= (n-i)/2; j-- {
			retSlice[n-(n-i)/2-1][j] = num
			num++
		}
		//写回型第一列
		for j := n - (n-i)/2 - 2; j > (n-i)/2; j-- {
			retSlice[j][(n-i)/2] = num
			num++
		}
	}
	return retSlice
}

func generateMatrix3(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	left := 0
	up := 0
	down := n - 1
	right := n - 1
	r:=make([][]int,n)
	for i:=0;i<n;i++{
		r[i]=make([]int,n)
	}
	k:=1
	for left <= right && up <= down {
		//上，要从0开始
		for i := left; i <= right; i++ {
			r[up][i]=k
			fmt.Println("up ",k)
			k++

		}
		//右，从1开始
		for j := up + 1; j <= down; j++ {
			r[j][right]=k
			fmt.Println("right ",k)
			k++

		}
		//下，左
		if left < right && up < down {
			//下方从 n-1开始，右下角-1
			for i := right - 1; i >= left; i-- {
				r[down][i]=k
				fmt.Println("down ",k)
				k++

			}
			//左边从最下面开始，也就是左下角-1
			for j := down - 1; j >= up+1; j-- {
				r[j][left]=k
				fmt.Println("left ",k)
				k++
			}
		}
		left++
		right--
		up++
		down--
	}
	return r
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix2(3))
	fmt.Println(generateMatrix3(4))
}
