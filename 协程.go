package main

import (
	//"fmt"
	"math"
	"runtime"
	"sync"
)
func sum1(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
func main()  {
	//简单说明一下用法，var 是声明了一个全局变量 wg，
	// 类型是sync.WaitGroup，wg.add(2) 是说我有2个goroutine需要执行，
	//wg.Done 相当于 wg.Add(-1) 意思就是我这个协程执行完了。
	// wg.Wait() 就是告诉主线程要等一下，等他们2个都执行完再退出。

	wg := new(sync.WaitGroup)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum1(id)
		}(i)
	}


	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			println("here",i)
			if i == 3 { runtime.Gosched() }
		}
	}()
	go func() {
		defer wg.Done()
		println("Hello, World!")
	}()
	wg.Wait()
}