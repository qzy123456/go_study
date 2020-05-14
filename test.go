package main

import (
   "fmt"
   "sync"
)
import "unsafe"
const (
   aa = "abc"
   bb=len(aa)
   cc=unsafe.Sizeof(aa)
   zero = 0.0
)
func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)
   println()
   println(aa, bb, cc)
   const str = ` 第一行
第二行
第三行
\r\n
`
   //fmt.Println(str)
   //x := 1
   //
   //{
   //
   //   x := 2
   //
   //   fmt.Print(x)
   //
   //}
   //fmt.Println(x)
   // 21
   //strs := []string{"one","two", "three"}
   //
   //
   //for _, s := range strs {
   //
   //   go func() {
   //
   //      time.Sleep(1 * time.Second)
   //
   //      fmt.Printf("%s ", s)
   //
   //   }()
   //
   //}
   //
   //time.Sleep(3 * time.Second)
   // three three three
   //s := NewSlice()
   //
   //defer s.Add(1).Add(2)
   //
   //s.Add(3)
   //132

   //fmt.Println(reflect.TypeOf(zero))
   //打印float64

   var number = 0
   var countGuard sync.Mutex
   go func(n *int){
      countGuard.Lock()
      // 在函数退出时解除锁定
      defer countGuard.Unlock()
      number++
   }(&number)


   go func(n *int){
      countGuard.Lock()
      // 在函数退出时解除锁定
      defer countGuard.Unlock()
      number++
   }(&number)

}



type Slices []int

func NewSlice() Slices {

   return make(Slices, 0)

}

func (s* Slices) Add(elem int) *Slices {

   *s = append(*s, elem)

   fmt.Print(elem)

   return s

}