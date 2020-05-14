package main

import (
	"fmt"
	"reflect"
)
func main() {

	// 声明一个空结构体
	type cat struct {
		Name string  `json:"type" name:"大师"`
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
	// 通过字段名, 找到字段类型Name信息
	if catType, ok := typeOfCat.FieldByName("Name"); ok {
		// 从
		fmt.Println(catType)
	}

	// 声明整型变量a并赋初值
	var a string = "1024"
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	// 获取interface{}类型的值, 通过类型断言转换
	var getA  = valueOfA.Interface()
	// 获取64位的值, 强制类型转换为int类型
	//var getA2 int = int(valueOfA.Int())
	fmt.Println(reflect.TypeOf(getA))
	var i interface{} = "77"
	value, ok := i.(int)
	if ok {
		fmt.Printf("类型匹配int:%d\n", value)
	} else {
		fmt.Println("类型不匹配int\n")
	}
	if value, ok := i.(int); ok {
		fmt.Println("类型匹配整型：%d\n", value)
	} else if value, ok := i.(string); ok {
		fmt.Printf("类型匹配字符串:%s\n", value)
	}

}