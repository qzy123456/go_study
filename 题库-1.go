package main

import (
	"fmt"
	"sync"
)

func main()  {
	pase_student()

	t := Teacher{}
	t.ShowA()  //showAshowB
	t.ShowB()  //teachershowB

	a := 1
	b := 2
	defer calc("1", a,calc("10", a, b))
	defer calc("2", a,calc("20", a, b))
	//首先，index "1"肯定是最后执行的，但是 index1中有一个是个函数，所以会被先调用
	//calc（'10'，a，b） 10 ，1，2，3
	//然后同理 index '2'中的函数  calc （'20'，a，b） 执行  20，1，2
	//执行 index '2'  calc（'2'，1，2） 2，1，2，3
	//执行 index '1'  calc（'1'，1，2） 1，1，2，3
	//10 1 2 3
	//20 1 2 3
	//2 1 3 4
	//1 1 3 4

	s := make([]int,5)
	s = append(s,1, 2, 3)
	fmt.Println(s)
	//[0 0 0 0 0 1 2 3] 一共8位,因为初始化已经是5个0了
	//如果改为
	ss := make([]int, 0)
	ss = append(ss, 1, 2, 3)
	fmt.Println(ss)
	//[1 2 3]
	var aa = make(map[string]int)
	aa["name"] = 18
	stru := UserAges{ages:aa}
	fmt.Println(stru.Get("11"))
	//下列报错是哪一行
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	//_ = y == y  //comparing uncomparable type []int
	f := F(5)
	defer func() {
		fmt.Println("第一个defer",f())  // 8
	}()
	defer fmt.Println("第二个defer",f()) // 6
	i := f()
	fmt.Println("defer结束",i) //7
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}
type Orange struct {
	Quantity int
}
func (o *Orange) Increase(n int) {
	o.Quantity += n }
func (o *Orange) Decrease(n int) {
	o.Quantity -= n }
func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity) }
func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func calc(index string, a, b int) int {
	ret := a+ b
	fmt.Println(index,a, b, ret)
	return ret
}

type student1 struct {
	Name string
	Age  int
}
func pase_student() {
	m := make(map[string]*student1)
	stus := []student1{
		{Name: "zhou",Age: 24},
		{Name: "li",Age: 23},
		{Name: "wang",Age: 22},
	}
	//错误，因为会使用同一个地址  所以map里面的值都是一样的
	for _,stu := range stus {
		m[stu.Name] =&stu
	}
	for k,v:=range m{
		println(k,"=>",v.Name)
	}
	//正确
	for i:=0;i<len(stus);i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k,v:=range m{
		println(k,"=>",v.Name)
	}
}

//继承
type People1 struct{}
func (p *People1) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("showB")
}
type Teacher struct {
	People1
}
func (t *Teacher) ShowB() {
	fmt.Println("teachershowB")
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}
func (ua *UserAges)Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string)int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

