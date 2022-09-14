package main

import (
	"sync"
)

var count = 5

func main() {

	wg := sync.WaitGroup{}
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1) //B和C可以无缓冲 make(chan struct{}),让A先执行就好了
	chanC := make(chan struct{}, 1)

	chanA <- struct{}{}
	wg.Add(3)

	go printA(&wg, chanA, chanB)
	go printB(&wg, chanB, chanC)
	go printC(&wg, chanC, chanA)
	wg.Wait()
}

func printA(wg *sync.WaitGroup, chanA, chanB chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanA
		println("A")
		chanB <- struct{}{}
	}
}

func printB(wg *sync.WaitGroup, chanB, chanC chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanB
		println("B")
		chanC <- struct{}{}
	}
}

func printC(wg *sync.WaitGroup, chanC, chanA chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanC
		println("C")
		chanA <- struct{}{}
	}
}
