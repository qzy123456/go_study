package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	str := "hello，朝阳!"
	str1 := []rune(str)
    //定义一个长度和字符串一样的chan
	sc := make(chan rune, len(str))
	//当前轮到谁打印的信号
	sigle := make(chan struct{})
    //往chan里面写数据
	for _, v := range str1 {
		sc <- v
	}

	close(sc)
	//开始
	sigle <- struct{}{}

	go func() {
		defer wg.Done()
		for {
			ball , ok := <- sigle
			if ok {
				pri, ok1 := <- sc
				if ok1 {
					fmt.Printf("go 1 : %c\n", pri)
				} else {
					close(sigle)//字符串打完，可以关闭了
					return
				}
			} else {
				return
			}
			sigle <- ball
		}
	}()

	go func() {
		defer wg.Done()
		for {
			ball , ok := <- sigle
			if ok {
				pri, ok1 := <- sc
				if ok1 {
					fmt.Printf("go 2 : %c\n", pri)
				} else {
					close(sigle)//字符串打完，可以关闭了
					return
				}
			} else {
				return
			}
			sigle <- ball
		}
	}()
	wg.Wait()
}