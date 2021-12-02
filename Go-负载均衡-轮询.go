
package main

import (
"fmt"
	"sync"
)

type RoundRobin struct {
	servers []string
	current int
	sync    sync.RWMutex
}

/**
获取下一个服务器
*/
func (R *RoundRobin) next() string {
	R.sync.Lock()
	defer R.sync.Unlock()
	R.current++
	R.current = R.current % len(R.servers) // 访问到最后一个服务器之后，重置会第一台。 5%5=0。
	return R.servers[R.current]
}

func main() {

	r := &RoundRobin{
		servers: []string{"192.168.10", "192.168.11", "192.168.12"},
		current: -1,
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("| %d | %s |\n", i + 1, r.next())
	}
}

