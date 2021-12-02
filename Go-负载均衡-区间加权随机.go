package main

import (
	"fmt"
	"math/rand"
)

type nodeWeight struct {
	Node   string
	Weight int
}

type nodeOffset struct {
	Node        string
	Weight      int
	OffsetStart int
	OffsetEnd   int
}

var (
	weightList = []nodeWeight{
		nodeWeight{
			Node:   "127.0.0.1",
			Weight: 5,
		},
		nodeWeight{
			Node:   "127.0.0.2",
			Weight: 30,
		},
		nodeWeight{
			Node:   "127.0.0.3",
			Weight: 2,
		},
		nodeWeight{
			Node:   "127.0.0.4",
			Weight: 1,
		},
		nodeWeight{
			Node:   "127.0.0.5",
			Weight: 8,
		},
	}

	offsetList = make([]nodeOffset, 0)

	totalWeight = 0

	res = make(map[string]int)

	forCount = 0
)

func Init() {
	for k, v := range weightList {
		tmp := nodeOffset{}

		if k == 0 {
			tmp = nodeOffset{
				Node:        v.Node,
				Weight:      v.Weight,
				OffsetStart: 0,
				OffsetEnd:   v.Weight,
			}
		} else {
			tmp = nodeOffset{
				Node:        v.Node,
				Weight:      v.Weight,
				OffsetStart: totalWeight + 1,
				OffsetEnd:   totalWeight + v.Weight,
			}
		}

		totalWeight = totalWeight + v.Weight
		offsetList = append(offsetList, tmp)
	}

}

func getNodeAddress() string {

	if totalWeight <= 0 {
		panic("address is empty")
	}

	idx := rand.Intn(totalWeight)

	node := ""
	for _, v := range offsetList {
		forCount = forCount + 1
		if idx >= v.OffsetStart && idx <= v.OffsetEnd {
			node = v.Node
			break
		}
	}

	if _, ok := res[node]; !ok {
		res[node] = 1
		return node
	}

	res[node] = res[node] + 1
	return node
}

func main() {
	Init()

	i := 1
	for {
		if i > 10000 {
			break
		}

		getNodeAddress()
		i++
	}
	fmt.Println(res)
	fmt.Println(forCount)
}