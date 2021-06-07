package main

import "fmt"
//贪婪匹配，无限买入卖出
func maxProit(price []int)int  {
	var max  = 0
	for i:=1;i<len(price) ;i++  {
		if price[i] > price[i-1]{
			max += price[i] - price[i-1]
		}
	}
	return max
}
func main() {

	price := []int{7,1,5,3,6,4}
	fmt.Println(maxProit(price))
}
