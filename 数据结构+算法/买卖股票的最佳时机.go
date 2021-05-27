package main

import "fmt"

//最大利润 = 第i天买入 - 最小买入
func maxProit2(price []int)int  {
	var min_input = price[0]
	var max_point  = 0
	for i:=1;i<len(price) ;i++  {
		min_input = min(min_input,price[i])
		max_point = max(max_point,price[i] - min_input)
	}
	return max_point
}
func max(a,b int)int  {
     if a>=b {
     	return a
	 }
     return b
}
func min(a,b int)int  {
	if a<= b{
		return a
	}
	return b
}
func main() {

	price := []int{7,1,5,3,6,4}
	fmt.Println(maxProit2(price))
}
