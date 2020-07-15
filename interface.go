package main
import ("fmt"
	//"reflect"
)
// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}
// 结构体类型
type Struct struct {
}
// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}
// 函数定义为类型
type FuncCaller func(interface{})
// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}
func main() {
	// 声明接口变量
	var invoker Invoker
	// 实例化结构体
	s := new(Struct)
	// 将实例化的结构体赋值到接口
	invoker = s
	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")
	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
	//匿名接⼝口可⽤用作变量类型，或结构成员。
	t := Tester{&Userr{1, "Tom"}}
	fmt.Println(t.s.String())
    //还可⽤用 switch 做批量类型判断，不⽀支持 fallthrough。
	var o interface{} = &Userr{1, "Tom1111"}
	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case *Userr:
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}

type Tester struct {
	s interface {
		String() string
	}
}
type Userr struct {
	id   int
	name string }
func (self *Userr) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}
