package main

import "fmt"

//暴力破解
func bf(n int)int  {
	sum := 0
	for i:=2;i<n;i++{
		if isPrime(i){
			sum++
		}
	}
	return sum
}

func isPrime(n int)bool  {
	for i:=2;i*i<=n;i++ {
		if n%i ==0{
			return  false
		}
	}
	return true
}

//埃筛法（埃式筛选法）
func eratostheness(n int)int{
	var isComposite = make([]bool,n)
	var count = 0
	for i:=2;i<n ;i++  {
		//如果是质数
		if !isComposite[i]{
			count++
			// i*i, ..., 4为合数
			for j:=i*i;j<n ;j+=i  {
				isComposite[j] = true
			}
		}
	}
    return count
}

func main() {
	fmt.Println(bf(9))
	fmt.Println(eratostheness(9))
}
