package main

import (
	"errors"
	"fmt"
)

func main()  {
    fmt.Println(Test1()) //1
    fmt.Println(Test2()) //2
    fmt.Println(Test3()) //4
	e1() //nil
	e2() //e2 defer err
	e3() //nil
}
// 测试1
func Test1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}
func Test2() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 2
}
func Test3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}
func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}