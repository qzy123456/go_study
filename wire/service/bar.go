// service/bar.go

package service

import "fmt"

type Bar interface {
	Bar()
}

type bar struct {
}

func (b *bar) Bar() {
	fmt.Println("bar")
}

func NewBar() Bar {
	return &bar{}
}