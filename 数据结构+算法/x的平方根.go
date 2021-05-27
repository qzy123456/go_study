package main

import "fmt"

//递归
func newton(n int)int  {
	if n==0{
		return 0
	}
	return int(sqrt1(float64(n),float64(n)))
}
func sqrt1(x,n float64) float64{
	res := (x+n/x)/2
	if res == x{
		return x
	}
	return sqrt1(res,n)
}

//循环
func newton2(x int)int  {
	if x ==0 {
		return 0
	}
	var n  = x
	for n*n>x{
		n = (n +x/n)/2
	}
	return int(n)
}
//二分法
func newton3(x int)int  {
	var left  = 0
	var right = x
	for left <= right {
		mid := (left + right) /2
		if mid * mid == x{
			return mid
		}else if mid*mid<x{
			left = mid + 1
		}else if mid*mid >x{
			right = mid - 1
		}
	}
	return right
}

func main() {
	fmt.Println(newton(10))
	fmt.Println(newton2(10))
	fmt.Println(newton3(10))
}
