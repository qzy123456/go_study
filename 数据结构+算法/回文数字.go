package main

import "fmt"

func isPalindrome(x int) bool {
	if x<0{
		return false
	}

	var result int
	for i:=x; i!=0; i=i/10{
		fmt.Println(i)
		tmp := i%10
		fmt.Println(tmp)
		result = result*10 + tmp
		fmt.Println(result)
	}

	return result==x
}

func main() {
	fmt.Println(isPalindrome(121))
}
