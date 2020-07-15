package main

import (
	"fmt"
)

// 定义一个接口USB,有两个方法
type USB interface {
	start()
	stop()
}

// 定义Phone的类型的结构体
type Phone struct {
	name string
}

// Phoen类型实现接口USB
func (p Phone) start() {
	fmt.Println(p.name, "start")
}
func (p Phone) stop() {
	fmt.Println(p.name, "stop")
}

// 定义Pad类型的结构体
type Pad struct {
	name string
}

// Pad实现接口USB
func (p Pad) start() {
	fmt.Println(p.name, "start")
}
func (p Pad) stop() {
	fmt.Println(p.name, "stop")
}

func main() {
	// 定义一个map变量iArr它的键是string类型,值是USB接口类型
	// 那么只有实现了USB接口的数据才能符合
	var iArr map[string]USB
	iArr = make(map[string]USB)
	// Phone结构体类型 实现了USB ,所以可以将Phone类型的结构体可以当做值赋给变量IArr
	iArr["iPhone X"] = Phone{"iPhonex"}
	// ipad结构体类型 实现了USB ,所以可以将ipad类型的结构体可以当做值赋给变量IArr
	iArr["IPad2018"] = Pad{"Ipad2020款"}
	// 调用方法
	iArr["IPad2018"].start()
	iArr["IPad2018"].stop()
	iArr["P30"] = Phone{"华为P30"}
	iArr["Mate20"] = Phone{"Mate 20 X"}
	iArr["P30"].stop()
	iArr["P30"].start()
}

