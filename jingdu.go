package main

import (
	"fmt"
	"math"
	"unsafe"
	"text/template"
)

const (
	s = "abc" //在常量组中，如不提供类型和初始化的值，那么视为与上一常量相同（第一个常量值不能为空）
	x
	// x = "abc"
)
const (
	a = "abc" //常量值还可以是len cap等编译期间可确定结果的函数返回值
	b = len(a)
	ccx = unsafe.Sizeof(b)
)

//关键字 iota 定义常量组中从 0 开始按⾏行计数的⾃自增枚举值。
const (
	Sunday    = iota // 0
	Monday           // 1，通常省略后续⾏行表达式。
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

//如果 iota ⾃自增被打断，须显式恢复
const (
	A = iota //0
	B        // 1
	C = "c"  //c
	D        // c，与上⼀一⾏行相同。
	E = iota // 4，显式恢复。注意计数包含了 C、D 两⾏行。
	F        // 5
)
// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}
func length(s string) int {
	println("call length.")
	return len(s)
}
func main() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	x := 0x100
	defer println(x)
	fmt.Println(x)
	i := 0
	_ = i
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"
	//要修改字符串，可先将其转换成 []rune 或 []byte，
	//[]rune用来处理utf-8 中文字符
	// 完成后再转换为 string。⽆无论哪种转 换，都会重新分配内存，并复制字节数组。
	us := []rune(u)
	us[1] = '话'
	println(string(us))
	//用for循环遍历字符串时，也有byte和rune两种格式
	s = "abc汉字"
	for i := 0; i < len(s); i++ {     //【】byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s {           //【】rune
		fmt.Printf("%c,", r)
	}
	fmt.Println("\n")
	//⽀支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
	//• 默认值 nil，没有 NULL 常量。
	//• 操作符 "&" 取变量地址，"*" 透过指针访问⺫⽬目标对象。
	//• 不⽀支持指针运算，不⽀支持 "->" 运算符，直接⽤用 "." 访问⺫⽬目标成员。
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a) // 直接⽤用指针访问⺫⽬目标对象成员，⽆无须转换。

	uu := uint32(32)
	ii := int32(1)
	fmt.Println(&uu, &ii)
	pp := &ii
	//pp = (*int32)(&uu)这中会报错，说是类型不同 不让转换
	pp = (*int32)(unsafe.Pointer(&uu))
	fmt.Println(pp)
	ttt := template.New("Foo")
	ppp := (*MyTemplate)(unsafe.Pointer(ttt))
	ppp.name = "Bar" // 关键在这里，突破私有成员
	fmt.Println(ppp, ttt)
	//不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。
	s = "abcd"
	for i, n := 0, length(s); i < n; i++ {
		println(i, s[i])
	}
L1:
	for x := 0; x < 3; x++ {
L2:
		for y := 0; y < 5; y++ {
			if y > 2 { continue L2 }
			if x > 1 { break L1 }
			print(x, ":", y, " ")
		}
		println() }
}
