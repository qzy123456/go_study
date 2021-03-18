package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	service := NewRequestLimitService(time.Second, 2)
	for true {
		hasToken := service.AddRequestCount()
		if hasToken {
			fmt.Println(time.Now())
		}
	}
}

type RequestLimiter struct {
	Interval time.Duration // 重新计数时间
	MaxCount int           // 最大计数
	Lock     sync.Mutex
	ReqCount int // 目前的请求数
}

// 非阻塞
func (reqLimiter *RequestLimiter) AddRequestCount() bool {
	reqLimiter.Lock.Lock()
	defer reqLimiter.Lock.Unlock()
	if reqLimiter.ReqCount < reqLimiter.MaxCount {
		reqLimiter.ReqCount += 1
		return true
	}
	return false
}
//初始化限流器，计时器
func NewRequestLimitService(interval time.Duration, maxCount int) *RequestLimiter {
	reqLimit := &RequestLimiter{
		Interval: interval,
		MaxCount: maxCount,
	}
	go func() {
		ticker := time.NewTicker(interval)
		for true {
			<-ticker.C
			reqLimit.Lock.Lock()
			fmt.Println("reset Count ...")
			reqLimit.ReqCount = 0
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}
