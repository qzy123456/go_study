package main

import "fmt"

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

func main() {
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, 3))
	fmt.Println(convert2(s, 3))
}
