package main

import "fmt"
//最多只能买2次
func maxProitThree(price []int)int  {
	if len(price) ==0 || len(price) <=2{
		return 0
	}
	var L2R = make([]int,len(price)) //左到右
	var R2L = make([]int,len(price)) //右到左

	minPrice := price[0]
	p1 :=0
	for i:=1;i<len(price) ;i++  {
		if price[i] - minPrice > p1{
			p1 = price[i] - minPrice
		}
		L2R[i] = p1
		if price[i] < minPrice{
			minPrice = price[i]
		}
	}
	maxPrice := price[len(price) -1]
	p2 :=0
	for i:=len(price) -2;i >=0 ;i--  {
		if maxPrice - price[i] > p2{
			p2 = maxPrice - price[i]
		}
		R2L[i] = p2
		if maxPrice > price[i]{
			maxPrice = price[i]
		}
	}
	k:=0
	//左右重合求最大
	for i:=0;i<len(price) ;i++  {
		if L2R[i] + R2L[i] > k{
			k = L2R[i] + R2L[i]
		}
	}
	return k
}
func main() {

	price := []int{1,2,3,4,5}
	fmt.Println(maxProitThree(price))
}
