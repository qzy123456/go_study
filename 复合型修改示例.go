package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()
	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)
	//dayin [100 200] [100 200]
	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))
	fmt.Println(v2.Interface())
	//[100 200 300 400 500]
	fmt.Println("----------------------")
	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()
	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(v.Interface(), m)
	today := time.Now().Format("2006-01-02")
	fmt.Println(today)
	//#string到int
	int1, _ := strconv.Atoi("123")
	fmt.Println(reflect.TypeOf(int1))
	//#string到int64
	int641, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(reflect.TypeOf(int641))
	//#int到string
	string1 := strconv.Itoa(int1)
	fmt.Println(reflect.TypeOf(string1))
	//#int64到string
	string1 = strconv.FormatInt(int641, 10)
	fmt.Println(reflect.TypeOf(string1))

	f, _ := strconv.ParseFloat(string1, 32)
	fmt.Println(reflect.TypeOf(f))
	//map[a:100 b:200] map[a:100 b:200]
	b := []byte(`{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
}`)

	var t interface{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t)

	//使用断言判断类型
	m1 := t.(map[string]interface{})
	for k, v := range m1 {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case bool:
			fmt.Println(k, "is bool", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
