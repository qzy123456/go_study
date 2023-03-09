package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan struct{})

	go func() {
		defer wg.Done()
		for i:=0;i<10 ;i++  {
			if i%2 == 0 {
				fmt.Println(i)
			}
			ch <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i:=0;i<10 ;i++  {
			<- ch
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}