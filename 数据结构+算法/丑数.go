package main

import "fmt"

//对n不断除以5,3,2，最后看是否能除尽（即最后等于1）
func isUgly(n int) bool{
    if n<=0{
    	return  false
	}
    if n ==1 {
    	return true
	}
    if n%5==0{
    	n/=5
	}
	if n%3==0{
		n/=3
	}
	if n%2==0{
		n/=2
	}

    return n==1
}

func main() {
	fmt.Println(isUgly(20))
}
