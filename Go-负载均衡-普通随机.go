
package main

import (
	"fmt"
	"math/rand"
)

type Random struct {
	servers []string
}

func (R *Random) next() string {
	return R.servers[rand.Intn(len(R.servers))]
}

func main()  {
	r := Random{
		servers: []string{"192.168.10.10", "192.168.10.11", "192.168.10.12"},
	}

	for i := 0; i < 10; i++ {
		fmt.Println(r.next())
	}
}