package main

import (
	"fmt"
	"sync"
	"net/url"
)

var hLock sync.Mutex;

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go testt(ch, i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	// url encode
	v := url.Values{}
	v.Add("msg", "此订单不存在或已经提交")
	body := v.Encode()
	fmt.Println(v)
	fmt.Println(body)
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m)
}
func testt(ch chan int, i int) {
	hLock.Lock()
	var f  = make([]byte,1)
	f[0] =byte(i)
	fmt.Println(f)
	ch <- i
	hLock.Unlock()
}