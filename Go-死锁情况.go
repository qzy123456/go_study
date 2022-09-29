package main

import (
	"sync"
	"time"
)

var l sync.RWMutex

func lockAndRead() { // 可读锁内使用可读锁
	l.RLock()
	defer l.RUnlock()

	l.RLock()
	defer l.RUnlock()
}

func lockAndRead1() { // 全局锁内使用全局锁
	l.Lock()
	defer l.Unlock()

	l.Lock()
	defer l.Unlock()
}

func lockAndRead2() { // 全局锁内使用可读锁
	l.Lock()
	defer l.Unlock() // 由于 defer 是栈式执行，所以这两个锁是嵌套结构

	l.RLock()
	defer l.RUnlock()
}

func lockAndRead3() { // 可读锁内使用全局锁
	l.RLock()
	defer l.RUnlock()

	l.Lock()
	defer l.Unlock()
}

	//RWMutex的使用主要事项
	//1、读锁的时候无需等待读锁的结束
	//2、读锁的时候要等待写锁的结束
	//3、写锁的时候要等待读锁的结束
	//4、写锁的时候要等待写锁的结束
    //谨防锁拷贝
    //type MyMutex struct {
	//	count int
	//	sync.Mutex
	//}
	//func main() {
	//	var mu MyMutex
	//	mu.Lock()
	//	var mu1 = mu  //panic
	//	mu.count++
	//	mu.Unlock()
	//	mu1.Lock()
	//	mu1.count++
	//	mu1.Unlock()
	//	fmt.Println(mu.count, mu1.count)
	//}
	//加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁

func main() {
	lockAndRead()
	time.Sleep(5 * time.Second)
}

