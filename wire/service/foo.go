// service/foo.go

package service

import "fmt"

type Foo interface {
	Foo()
}

type foo struct {
}

func (f *foo) Foo() {
	fmt.Println("foo")
}

func NewFoo() Foo {
	return &foo{}
}