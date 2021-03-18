package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //桶上次放token的时间戳 s
	lock         sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().Unix()
	//根据速率，时间 计算当前桶中前挡的数量，也就是1秒产生多少个
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate
    //判断容量超出
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	//更新时间
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		return false
	}
}

//令牌桶
func main() {
	s := &TokenBucket{
		rate:         1, //1秒1个
		capacity:     2, //容量2个
		tokens:       1, //当前1个
		lastTokenSec: time.Now().Unix(),
	}
	for true {
		if s.Allow() {
			fmt.Println("done")
		}
	}
}
