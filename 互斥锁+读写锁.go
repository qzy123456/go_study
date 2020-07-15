package main
import (
	"fmt"
	"sync"
)
var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
	// 从 sync.Mutex 互斥锁改为 sync.RWMutex 读写互斥锁。
	countGuard2 sync.RWMutex
)
func GetCount() int {
	// 锁定
	countGuard2.RLock()
	// 在函数退出时解除锁定
	defer countGuard2.RUnlock()
	return count
}
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}
func main() {
	// 可以进行并发安全的设置
	SetCount(1)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}