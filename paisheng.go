package main

import (
	"fmt"
	"time"
	"context"
)

type Cat struct {
	Color string
	Name  string
}
type BlackCat struct {
	Cat  // 嵌入Cat, 类似于派生   利用继承机制，新的类可以从已有的类中派生。那些用于派生的类称为这些特别派生出的类的“基类”。
}
// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}
///////////////////////////////////////////////////////////
//向结构体添加方法
type Bag struct {
	items []int
}
//插入到结构体的方法-1
func Insert(b *Bag,itemId int)  {
	b.items = append(b.items,itemId)
}
//插入到结构体的方法-2
//(b*Bag) 表示接收器，即 Insert 作用的对象实例。
func (b *Bag) Inserts(itemId int)  {
	b.items = append(b.items,itemId)
}
///////////////////////////////////////////////////////////
//定义属性结构，属性值

type Property struct {
	value int //属性值
}
//设置属性值
func (p *Property) setValue(v int)  {
	p.value = v
}
//得到属性值
func  (p *Property) values() int  {
	return p.value
}
///////////////////////////////////////////////
//非指针类型的接收器
type Point struct {
	x int
	y int
}
//非指针的接收器的add方法
func (p Point) Add (other Point)Point  {
	//成员值互加之后返回新的结构
	return Point{p.x + other.x , other.y+p.y}
}
func main()  {
	//派生🐱lei
	cat1 := NewCat("cat1")
	cat2 := NewBlackCat("black cat2")
	fmt.Println(cat1,cat2,"\n")
	////////////////////////////////////////////
	//实例化结构体的方法
	bag := new(Bag)
	Insert(bag,111)
	//接收器调用
	bag.Inserts(222)
	fmt.Println(bag,"\n")
	/////////////////////////////////////////////
	//实例化属性值，赋值，得到
	p := new(Property)
	p.setValue(21221)
	fmt.Println(p.values(),"\n")
	//////////////////////////////////////////////
   //非指针的接收器
   p1 := Point{1,1}
   p2 := Point{2,2}
   res := p1.Add(p2)
   fmt.Println(res,"\n")
	d := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancel action function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}