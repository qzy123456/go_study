package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	//是否需要转正负
	neg, res := false, 0
	//除数小于0
	if dividend < 0 {
		dividend, neg = -dividend, !neg
	}
	//被除数小于0
	if divisor < 0 {
		divisor, neg = -divisor, !neg
	}
	sum, cnt, multi := 0, 1, divisor
	for cnt >= 1 {
		if sum+multi <= dividend {
			sum, res, cnt, multi = sum+multi, res+cnt, cnt<<1, multi<<1
		} else {
			cnt, multi = cnt>>1, multi>>1
		}
	}

	if neg {
		res = -res
	}
	if res > (1<<31)-1 {
		res = (1 << 31) - 1
	}
	return res
}

func divide2(dividend int, divisor int) int {
	res,flag:= 0,true
	if (dividend ^ divisor)<0{
		flag = false
	}
	if dividend<0{
		dividend = -dividend
	}
	if divisor<0{
		divisor = -divisor
	}
	//位运算
	for i:=31;i>=0;i--{
		if (dividend>>i) >= divisor{
			res += 1<<i
			dividend -= divisor<<i
		}
	}
	if !flag{
		res=-res
	}
	if res>math.MaxInt32{
		res = math.MaxInt32
	}
	return res
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide2(10, 3))
}
