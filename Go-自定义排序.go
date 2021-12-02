package main

import (
	"fmt"
	utils "go-study/models"
)

func main() {

	// int 集合
	s := utils.NewIntSet()

	// 添加数据
	s.Add(5, 2, 4, 3, 5, 6, 7)

	// 去重后的值
	fmt.Println(s.List())

	// 排序后的值
	fmt.Println(s.SortList())



	// string 集合
	s2 := utils.NewStringSet()

	// 添加数据
	s2.Add("wen", "jian", "bao", "study", "goalng", "bao", "jian")

	// 去重后的值
	fmt.Println(s2.List())

	// 排序后的值
	fmt.Println(s2.SortList())
}