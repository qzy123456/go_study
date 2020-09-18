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
		for elem, value := range set.s{
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
	s:[]interface{}{"1", "2"},
}
v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
}
