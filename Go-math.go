package main

import (
	"fmt"
	"math"
)

const BASE  = 0.5
func main(){
	var num float64
	var xx  float64
   for{
   	fmt.Println("请输入一个数字，并回车:")
   	fmt.Scanln(&num)
   	xx = num * 500
   	xx = math.Pow(xx,BASE) + 200
   	fmt.Println( math.Round(xx))
   }


}
