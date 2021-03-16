// 滑动窗口
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	service := NewRequestLimitServices(time.Second, 2, 1)
	for true {
		hasToken := service.AddRequestCount()
		if hasToken {
			fmt.Println(time.Now())
		}
	}
}

type WindowLimiter struct {
	Interval    time.Duration // 总计数时间
	WinCount    []int         // 每个窗口的访问数量
	TicketSize  int           // 窗口最大容量
	TicketCount int           // 窗口数量
	Lock        sync.Mutex
	CurIndex    int // 目前使用哪个窗口
}

func (reqLimiter *WindowLimiter) IsAvailable() bool {
	reqLimiter.Lock.Lock()
	defer reqLimiter.Lock.Unlock()
	return reqLimiter.WinCount[reqLimiter.CurIndex] < reqLimiter.TicketSize
}

// 非阻塞
func (reqLimiter *WindowLimiter) AddRequestCount() bool {
	reqLimiter.Lock.Lock()
	defer reqLimiter.Lock.Unlock()
	if reqLimiter.WinCount[reqLimiter.CurIndex] < reqLimiter.TicketSize {
		reqLimiter.WinCount[reqLimiter.CurIndex]++
		return true
	}
	return false
}

func NewRequestLimitServices(interval time.Duration, ticketCount int, ticketSize int) *WindowLimiter {
	reqLimit := &WindowLimiter{
		Interval:    interval,
		WinCount:    make([]int, ticketCount, ticketCount),
		TicketSize:  ticketSize,
		TicketCount: ticketCount,
		CurIndex:    0,
	}
	go func() {
		ticker := time.NewTicker(time.Duration(interval.Nanoseconds() / int64(ticketCount)))
		for true {
			<-ticker.C
			reqLimit.Lock.Lock()
			reqLimit.CurIndex = (reqLimit.CurIndex + 1) % reqLimit.TicketCount
			reqLimit.WinCount[reqLimit.CurIndex] = 0
			fmt.Println("reset Count ...")
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}
