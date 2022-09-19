package main

import "fmt"

func main() {
	fmt.Println(f1())  //1
	fmt.Println(f2())  //5
	fmt.Println(f3())  //1
	f44()
	//a defer2: 2
	//a defer1: 3
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func f44() {
	i := 0
	//压栈，
	defer func() {//defer1
		i++//2+1
		fmt.Println("a defer1:", i)//i=3
	}()
	//defer压栈，先计算这个
	defer func() {//defer2
		i++//1+1
		fmt.Println("a defer2:", i)//i=2
	}()
	i++//i=1
}
