package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UserUser struct {
	Id   int
	Name string
	Age  int
}

func (u UserUser) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u UserUser) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
	user := UserUser{1, "Allen.Wu", 25}

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法,
	//PS: 无参数接口也要给参数，不然会报错。很奇葩
	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value,0)
	methodValue.Call(args)


	//对应下面的逻辑哦
	o := order{
		ordId:      456,
		customerId: 56,
	}
	query(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	query(e)

	// Split() 字符串分割
	slice2 := strings.Split("-hello-world-haha-","-")  // 返回字符串切片 []string
	fmt.Println(slice2)  // [ hello world haha ]  (两端有两个空元素)
	fmt.Println(reflect.TypeOf(slice2))  // 5
	cc :=strings.Replace(strings.Trim(fmt.Sprint(slice2), "[ ]"), " ", ",", -1)

	fmt.Println(cc)


}

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}
func query(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		sql := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		var dou string  = ""
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:

				if i == 0 {
					query(query)
					sql += fmt.Sprintf("%d,%s",  v.Field(i).Int(),dou )
				} else {
					sql += fmt.Sprintf(" %d,%s", v.Field(i).Int(),dou)
				}

			case reflect.String:
				if i == 0 {
					sql += fmt.Sprintf("%s,%s", v.Field(i).String(),dou)
				} else {
					sql += fmt.Sprintf("%s,%s", v.Field(i).String(),dou)
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		sql = fmt.Sprintf("%s",strings.TrimRight(sql,","))
		sql = fmt.Sprintf("%s)", sql)
		fmt.Println(sql)
		return

	}
	fmt.Println("unsupported type")
}