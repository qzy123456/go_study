package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			//编译错误，err.Error undefined (type interface {} is interface with no methods)
			//fmt.Println("捕获异常:", err.Error())
			fmt.Println("捕获异常:", fmt.Errorf("%v", err).Error())
		}
	}()
	panic("a")
}
