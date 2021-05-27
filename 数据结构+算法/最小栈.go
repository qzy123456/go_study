package main

import "fmt"

//切片实现
type MinStack struct {
	stack []int
	minIndex int
	length int
}

//初始化
func initialize()MinStack  {
	return MinStack{}
}

func (this *MinStack) Push(val int)  {
	this.stack =append(this.stack,val)
	this.minIndex = this.CalculateMinIndex()
	this.length++
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack) -1]
	this.minIndex = this.CalculateMinIndex()
	this.length--
}

func (this *MinStack) Top()int  {
	return this.stack[this.length -1]
}
func (this *MinStack) GetMin()int  {
	return this.stack[this.minIndex]
}
func (this *MinStack) CalculateMinIndex() int {
	var min_index  = 0
	for i:=1;i<len(this.stack);i++ {
		if this.stack[i] < this.stack[min_index]{
			min_index = i
		}
	}
	return min_index
}

func main()  {
	stac := initialize()
	stac.Push(1)
	stac.Push(2)
	stac.Push(-1)
	fmt.Println(stac.GetMin()) //-1
	stac.Pop()
	fmt.Println(stac.stack)   //[1 2]
	fmt.Println(stac.Top())    //2
}