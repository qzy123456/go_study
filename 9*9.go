package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
/*
   #include <stdio.h>
   #include <stdlib.h>
   void hello() {
        printf("Hello, World!\n");
} */
import "C"
var	App int
var	Uuid int

type TT struct {
	App int
	Uuid int
}
//字段标签可以实现单元数据编程，比如标记ORM MODEL属性
type User struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}
//import "fmt"
func main() {
	data := make(chan int)	// 数据交换队列
	//data := make(chan int，3)	//缓冲区可以3个元素
	exit := make(chan bool)  // 退出通知

	go func() {
		for d := range data {	// 从队列迭代接收数据，直到 close 。
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true	// 发出退出通知。
	}()
	data <- 1		// 发送数据。
	data <- 2
	data <- 3
	data <- 4      //若缓冲区设置了大小，则会堵塞
	close(data)		// 关闭队列。
	fmt.Println("send over.")
	<-exit			// 等待退出通知。
	//slice
	x := []int{0,1,2,3,4,5,6,7}
	y := x[1:3:6]
	fmt.Println(y)
	//调用C语言的逻辑
	//C.hello()
	var u User
	t := reflect.TypeOf(u)
	f,_ := t.FieldByName("Name")
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
	uu := User{"Jack", 23}
	v := reflect.ValueOf(uu)
	p := reflect.ValueOf(&uu)
	p.Elem().FieldByName("Name").SetString("dsdsd")
	p.Elem().FieldByName("Age").SetInt(11)
	fmt.Println(v.CanSet(), v.FieldByName("Name").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Name").CanSet())
	fmt.Println(p.Elem().FieldByName("Name"))
	fmt.Println(p.Elem().FieldByName("Age"))
	//除了上面的方面 p.Elem().FieldByName("Name").SetString("dsdsd")
	//	p.Elem().FieldByName("Age").SetInt(11)
	// 还能使用下面的指针方法
	ff := p.Elem().FieldByName("Age")
	fmt.Println(ff.CanSet())
	// 判断是否能获取地址。
	 if ff.CanAddr() {
	   agee := (*int)(unsafe.Pointer(ff.UnsafeAddr()))
	    // 等同
		 // age := (*int)(unsafe.Pointer(f.Addr().Pointer()))
	    //其他字段的更改要重新判断生成 类似 ff := p.Elem().FieldByName("Name") 需要ff
	    //namee := (*string)(unsafe.Pointer(ff.UnsafeAddr()))
	  // 等同
	    *agee = 88
	    //*namee = "new name"
     }
	fmt.Println(uu)
}