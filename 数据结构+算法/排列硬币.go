package main

import "fmt"

//总共有n枚硬币，排成阶梯形状，满足第K行有K个硬币。
//求形成的完整阶梯行的总行数。
//例如n=5时，
//1
//1 1
//1 1
//此时完整的阶梯行的总行数为2，因为第3行未排满。

func arrangeCoins(n int)int  {
	for i:=1;i<=n ;i++  {
		n -= i
		if n <= i{
			return i
		}
	}
	return 0
}

//二分
func arrangeCoins2(n int)int  {
	var start = 0
	var end =  n / 2 + 1
	var mid = end
	for end - start > 1 {
		mid = (start + end) >> 1   //右移 处以2，向下取整
		if mid * (mid + 1) ==  2 * n {
			return mid
		} else if mid * (mid + 1) <=  2 * n {
			start = mid
		} else {
			end = mid
		}
	}
	if end * (end + 1) ==  2 * n{
		return end
	}
	return   start
}
//牛顿迭代 ---不懂
func arrangeCoins3(n int)int  {
	 if n == 0{
	 	return 0
	 }
	 return int(sqrt(float64(n),float64(n)))
}
func sqrt(x,n float64)float64  {
	res := (x+(2*n-x)/x)/2
	if res == x{
		return x
	}
	return sqrt(res,n)
}

func main() {
	fmt.Println(arrangeCoins(10))
	fmt.Println(arrangeCoins2(10))
	fmt.Println(arrangeCoins3(10))
}