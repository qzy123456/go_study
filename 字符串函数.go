package main

import (
	"fmt"
	"reflect"
)

type Admin struct {
	name string
	title string
	age int
}

func main()  {
	var u Admin
	t:= reflect.TypeOf(u)

	for i,n := 0,t.NumField();i<n;i++ {
		f:=t.Field(i)
		fmt.Println(f.Name,f.Type)
	}

	uu := new(Admin)
	tt := reflect.TypeOf(uu)
	if tt.Kind() == reflect.Ptr {
		//tt.Kind(),reflect.Ptr   打印出ptr
		tt = tt.Elem() }
	for i, n := 0, tt.NumField(); i < n; i++ {
		f := tt.Field(i)
		fmt.Println(f.Name, f.Type)
	}
	fmt.Println("**********************")
	//可直接⽤用名称或序号访问字段，包括⽤用多级序号访问嵌⼊入字段成员。
	t1 := reflect.TypeOf(u)
	f, _ := t1.FieldByName("title")
	fmt.Println(f.Name) //title
	f, _ = t1.FieldByName("name")
	fmt.Println(f.Name) //name
	// 访问嵌⼊入字段。
	f, _ = t1.FieldByName("age") // 直接访问嵌⼊入字段成员，会⾃自动深度查找。(struct里面嵌套struct)
	 fmt.Println(f.Name) //age
//////////////////////////
	t2 := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t2)  //int
	//除 struct，其他复合类型 array、slice、map 取值⽰示例。
	v3 := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v3.Len(); i < n; i++ {
		fmt.Println(v3.Index(i).Int())
	}
	fmt.Println("---------------------------")
	var aa = make(map[string]interface{})
	aa["a"] = int(1)
	aa["b"] =int(2)
	fmt.Println(aa)
	v4 := reflect.ValueOf(aa)
	for _, k := range v4.MapKeys() {
		fmt.Println(k.String(), v4.MapIndex(k).Interface())
	}
}