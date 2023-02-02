package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	Name  string
	Age   int
	Arr   [2]bool
	ptr   *int
	slice []int
	map1  map[string]string
}
type Tx struct {
	Name  string
	Age   int
}

func main() {
	t1 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}
	t2 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	// 报错 实例不能比较 Invalid operation: t1 == t2 (operator == not defined on T1)
	// fmt.Println(t1 == t2)
	// 指针可以比较
	fmt.Println(&t1 == &t2) // false

	t3 := &T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	t4 := &T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	fmt.Println(t3 == t4)                  // false
	fmt.Println(reflect.DeepEqual(t3, t4)) // true
	fmt.Printf("%p, %p \n", t3, t4)        // 0xc000046050, 0xc0000460a0
	fmt.Printf("%p, %p \n", &t3, &t4)      // 0xc000006030, 0xc000006038

	// 前面加*，表示指针指向的值，即结构体实例，不能用==
	// Invalid operation: *t3 == *t4 (operator == not defined on T1)
	// fmt.Println(*t3 == *t4)

	t5 := t3
	fmt.Println(t3 == t5)                  // true
	fmt.Println(reflect.DeepEqual(t3, t5)) // true
	fmt.Printf("%p, %p \n", t3, t5)        // 0xc000046050, 0xc000046050
	fmt.Printf("%p, %p \n", &t3, &t5)      // 0xc000006030, 0xc000006040
	//结构体内的数据都可以比较
	t11 := Tx{
		Name:  "yxc",
		Age:   1,
	}
	t21 := Tx{
		Name:  "yxc",
		Age:   1,
	}
	fmt.Println(t11 == t21) //true
}