package main

import "fmt"

type Studentt struct {
	name string
	age int
}

func (s *Studentt)SetName(name string) (*Studentt) {
	s.name = name
	return s
}

func (s *Studentt)SetAge(age int) (*Studentt) {
	s.age = age
	return s
}

func (d *Studentt) Run() {
	println(d.name+"'s","age is",d.age)
}
func main() {
	s := Studentt{}
	s.SetName("li").SetAge(18).Run()
	fmt.Println(s)
}
