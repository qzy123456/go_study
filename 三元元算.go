package main

import (
	"fmt"
	"github.com/ymzuiku/hit"
	 //."github.com/ymzuiku/hit"  当我们用.的时候，证明我们不需要包的名称 例如可以直接省略  hit.If  ==> If
	"log"
)

func main(){
	//下面这些默认不成立，也就是默认是false
	//nil
	//0
	//error
	//""
	//"false"
	//"0"
	//If(a, b) 类似 a && b   if a == 成立, return b, else return nil
	//Or(a, b) 类似 a || b   if a == 成立, return a, else return b
	value1 := hit.If(20 > 5, "ok", "cancel")
	log.Println(value1) // ok

	value2 := hit.If("test", "ok", "cancel")
	log.Println(value2) // ok

	value3 := hit.If("", "ok", "cancel")
	log.Println(value3) // cancel

	value4 := hit.If("false", "ok", "cancel")
	log.Println(value4) // cancel

	value5 := hit.If(5, "ok", "cancel")
	log.Println(value5) // ok

	value6 := hit.Or(0, "ok", "cancel")
	log.Println(value6) // cancel

	value7 := hit.If(nil, "ok")
	log.Println(value7) // nil

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	fmt.Println(s1)
	s2 := s1[2:6:7]
	fmt.Println(s2)
	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)

}