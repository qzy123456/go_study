package main

import "fmt"

func main()  {
   //向 slice 尾部添加数据，返回新的 slice 对象。
	s := make([]int, 0, 5)  //0 leng 5 map
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)

    //可直接修改 struct array/slice 成员。
    //ps 【3】这个长度还是要给，不然下面的下标赋值  会出错，说超出range index
	dd := [3]struct {
		x int
	}{}
	ss := dd[:]
	dd[1].x = 10
	ss[2].x = 20
	fmt.Println(dd)
	fmt.Printf("%p, %p\n", &dd, &dd[0])
	//⼀一旦超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s1 := data[:2:3]
	s1 = append(s1, 100, 200) // ⼀一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s1, data)// 重新分配底层数组，与原数组⽆无关。
	fmt.Println(&s1[0], &data[0]) // ⽐比对底层数组起始指针。
    //通常以 2 倍容量重新分配底层数组。在⼤大批量添加数据时，建议⼀一次性分配⾜足够⼤大的空 间，
    // 以减少内存分配和数据复制开销。
    // 或初始化⾜足够⻓长的 len 属性，改⽤用索引号进⾏行操 作。
    // 及时释放不再使⽤用的 slice 对象，避免持有过期数组，造成 GC ⽆无法回收。
	s22 := make([]int, 0, 1)
	c := cap(s22)
	for ii := 0; ii < 50; ii++ {
		s22 = append(s22, ii)
		if n := cap(s22); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c=n
		}
	}
//打印出
//cap: 1 -> 2
//cap: 2 -> 4
//cap: 4 -> 8
//cap: 8 -> 16
//cap: 16 -> 32
//cap: 32 -> 64
//*************************//copy*********************************************//
//函数 copy 在两个 slice 间复制数据，复制⻓长度以 len ⼩小的为准。两个 slice 可指向同⼀底层数组，允许元素区间重叠。
	data11 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s11 := data11[8:] //8 9
	s211 := data11[:5] // 0,1,2,3,4
	//fmt.Println(s211)
	//将第二个slice里的元素拷贝到第一个slice里，拷贝的长度为两个slice中长度较小的长度值
	copy(s11,s211 )          // dst:s2, src:s
	fmt.Println(s211)
	fmt.Println(s11)
	fmt.Println(data11)
	//[0 1 2 3 4]
	//[0 1]
	//[0 1 2 3 4 5 6 7 0 1]













}
