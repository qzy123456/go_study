package main

import (
	"fmt"
	"math"
)

const abortIndex = math.MaxInt8 -1

// 中间件结构体
type Context struct {
	Handlers []func(c *Context) //需要执行的方法切片
	index    int8               //偏移量
}

// 注册方法到中间件
func (this *Context) Use(fs ...func(c *Context)) {
	for _,f :=range fs{
		this.Handlers = append(this.Handlers, f)
	}
}

// 执行下一个方法
func (this *Context) Next() {
	if this.index < int8(len(this.Handlers)) {
		this.index++
		this.Handlers[this.index](this)
	}
}

// 执行handlers里面的第一个方法
func (this *Context) Run() {
	this.Handlers[0](this)
}

func (this *Context) GET(path string, f func(c *Context)) {
	this.Handlers = append(this.Handlers, f)
}

// 终止
func (this *Context) Abort() {
	this.index = abortIndex
}

//主方法
func main() {
	c := &Context{}
	c.Use(MiddlewareOne(),MiddlewareTwo())
	c.GET("/get", func(c *Context) {
		fmt.Println("执行具体方法")
	})
	c.Run()
}

//中间件一
func MiddlewareOne() func(c *Context) {
	return func(c *Context) {
		fmt.Println("MiddlerOne Start")
		c.Next()
		fmt.Println("MiddlerOne End")
	}
}
//中间件二
func MiddlewareTwo() func(c *Context) {
	return func(c *Context) {
		fmt.Println("MiddlerTwo Start")
		c.Next()
		fmt.Println("MiddlerTwo End")
	}
}
