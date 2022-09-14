package main

import (
	"fmt"
	"sync"
)
func main() {
	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
	nums := 5
	var wg sync.WaitGroup
	wg.Add(3)
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= nums; i++ {
			<- ch1
			fmt.Println(s)
			ch2 <- struct{}{}
		}
		<- ch1
	}("A")

	go func(s string) {
		defer wg.Done()
		for i := 1; i <= nums; i++ {
			<- ch2
			fmt.Println(s)
			ch3 <- struct{}{}
		}
	}("B")

	go func(s string) {
		defer wg.Done()
		for i := 1; i <= nums; i++ {
			<- ch3
			fmt.Println(s)
			ch1 <- struct{}{}
		}
	}("C")
    //让ch1先执行
	ch1 <- struct{}{}
	wg.Wait()
}
