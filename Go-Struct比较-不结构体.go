package main

import "fmt"

type T2 struct {
	Name  string
	Age   int
	Arr   [2]bool
	ptr   *int
	//maps  map[string]int
}

type T3 struct {
	Name  string
	Age   int
	Arr   [2]bool
	ptr   *int
	//maps  map[string]int
}

func main() {

	var ss1 T2
	var ss2 T3
	// Cannot use 'ss2' (type T3) as type T2 in assignment
	//ss1 = ss2
	ss3 := T2(ss2)
	fmt.Println(ss3==ss1) // true
	//TODO 但是struct里面有不可比较的值的时候，这个时候整个struct都不可以比较了
}