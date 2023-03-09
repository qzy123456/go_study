package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	str := "hello,world!"
	str1 := []byte(str)

	sc := make(chan byte, len(str))
	count := make(chan int)

	for _, v := range str1 {
		sc <- v
	}

	close(sc)

	go func() {
		defer wg.Done()
		for {
			ball , ok := <- count
			if ok {
				pri, ok1 := <- sc
				if ok1 {
					fmt.Printf("go 1 : %c\n", pri)
				} else {
					close(count)
					return
				}
				count <- ball
			} else {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			ball , ok := <- count
			if ok {
				pri, ok1 := <- sc
				if ok1 {
					fmt.Printf("go 2 : %c\n", pri)
				} else {
					close(count)
					return
				}
			} else {
				return
			}
			count <- ball
		}
	}()

	count <- -1
	wg.Wait()
}