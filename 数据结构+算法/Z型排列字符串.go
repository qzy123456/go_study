package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	m := make([][]byte, numRows)
	j := 0
	for i := 0; i < len(s); i++ {
		if j < numRows {
			m[j] = append(m[j], s[i])
		} else {
			m[numRows-2-j%numRows] = append(m[numRows-2-j%numRows], s[i])
		}
		j++
		j = j % (2*numRows - 2)
	}

	var r []byte
	for i := 0; i < numRows; i++ {
		r = append(r, m[i]...)
	}
	return string(r)
}

func convert2(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}
//1、找规律
//Z 字形很容易找到规律，假如 numRows=3 ，那么周期为4；假如 numRows=4 ，那么周期为6。
//
//那么第一行的字符索引为 2 numRows -2
//
//最后一行的字符索引为 （2 numRows -2）+ numRows-1
//
//内部的行 i 的字符索引为 (2 numRows -2)+i 和 (2 numRows -2)-i
func convert3(s string, numRows int) string {
	if 1 == numRows {
		return s
	}
	var ret string
	var T = 2 * numRows - 2 // 周期
	fmt.Println("周期",T)
	for i := 0; i < numRows; i++ {
		for j := 0; i + j < len(s); j += T {
			ret += s[i+j:i+j+1]
			fmt.Println("字符",ret)
			if i != 0 && i != numRows-1 && j+T-i < len(s) {
				ret += s[j+T-i:j+T-i+1]
				fmt.Println("字符2",ret)
			}
		}
	}
	return string(ret)
}
func convert4(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	//t是周期
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

func main() {
	s := "123456789"
	fmt.Println(convert(s, 3))
	fmt.Println(convert2(s, 3))
	fmt.Println(convert3(s, 3))
	fmt.Println(convert4(s, 3))
}
