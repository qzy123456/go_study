package main

import (
	"fmt"
	"sync"
) //下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！   //Iter: 0 (0x10a4ea0,0x10df6e0)
	//ch := make(chan interface{}, len(set.s))    //Iter: 0 (0x10a4ea0,0x10df6e0) Iter: 1 (0x10a4ea0,0x10df6e0)
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}
func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
	//当 i 的值为 0、128 是会发⽣相等情况，注意 byte 是 uint8 的别名
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
		}
		if m == -m {
			count++
		}
	}
	fmt.Println(count)
	d1 := data{"one"}
	d1.print()
	//var in printer = data{"two"} //data没有实现printer，&data才实现
	var in printer = &data{"two"}
	in.print()
	a := [3]int{0, 1, 2}
	s := a[1:2]  //1
	s[0] = 11
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21
	fmt.Println(a) //[0 11 12]
	fmt.Println(s) //[21 12 13]
	fmt.Println(doubleScore(0))     //0
	fmt.Println(doubleScore(20.0))  //40
	fmt.Println(doubleScore(50.0))  //50
}

type data struct {
	name string
}
func (p *data) print() {
	fmt.Println("name:", p.name) }
type printer interface {
	print()
}
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}