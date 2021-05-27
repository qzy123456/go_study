package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack struct {
	min   int
	stack list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64,
		}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}
	this.stack.PushBack(val)
}

func (this *MinStack) Pop() {
	if this.min == this.stack.Back().Value.(int) {
		this.stack.Remove(this.stack.Back())
		this.min = math.MaxInt64
		for e := this.stack.Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			if val < this.min {
				this.min = val
			}
		}
	} else {
		this.stack.Remove(this.stack.Back())
	}

}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	stac := Constructor()
	stac.Push(1)
	stac.Push(2)
	stac.Push(-1)
	fmt.Println(stac.GetMin()) //-1
	stac.Pop()
	fmt.Println(stac.stack) //{{0xc000082000 0xc000082030 <nil> <nil>} 2}
	fmt.Println(stac.Top()) //2
}
