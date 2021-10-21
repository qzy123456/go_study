package main

import "fmt"

//要求计算 Pow(x, n)
//这一题用递归的方式，不断的将 n 2 分下去。注意 n 的正负数，n 的奇偶性。
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x

}

func main() {
	x := 2.0
	n := 2
	fmt.Println(myPow(x,n))
	fmt.Println(myPow(2.000,-1))
}