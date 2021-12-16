package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func main() {
	// 生成 100 个 [-10, 10) 范围的安全随机数。
	for i := 0; i < 100; i++ {
		ret := RangeRand(-10, 10)
		fmt.Println(ret)
	}
}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
