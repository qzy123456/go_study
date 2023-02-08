package main

import (
	"fmt"
	"math/rand"
)

var (
	nodeWeight = map[string]int{
		"127.0.0.1": 5,
		"127.0.0.2": 2,
		"127.0.0.3": 2,
		"127.0.0.4": 1,
	}

	nodeAddress = make([]string, 0)
)

func getNodeAddress() string {
	for k, v := range nodeWeight {
		i := 1
		for {
			if i > v {
				break
			}
			nodeAddress = append(nodeAddress, k)
			i++
		}
	}

	idx := rand.Intn(len(nodeAddress))
	address := nodeAddress[idx]

	return address
}

func main() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(getNodeAddress())
		i++
	}
}