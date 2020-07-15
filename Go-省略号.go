package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)
var validate = validator.New()
func validates(res RegisterReq)error  {
	err := validate.Struct(res)
	if err != nil{

		return  err
	}
	return  nil
}

type RegisterReq struct {
	UserName string `validate:"gt=0"`
	PasswordNew string `validate:"gt=0"`
	PasswordRepeat string `validate:"gt=0"`
	Email string `validate:"gt=0,email"`
}
//变长的函数参数
func f1(parms ...int) {
	for i, v := range parms {
		fmt.Printf("%v %v\n", i, v)
	}
}
//变长的函数参数
func f2(parms []int) {
	for i, v := range parms {
		fmt.Printf("f2:%v %v\n", i, v)
	}
}
func main()  {
	q := [...]int{1,2,3}
	fmt.Printf("%T\n",q) //"[3]int"

	var arr1 []int
	arr2 := []int{1,2,3}
	arr1 = append(arr1,0)
	arr1 = append(arr1,arr2...)	 //arr2... 将切片arr2打散成 ==> arr1 = append(arr1,1,2,3)
	fmt.Printf("%v\n",arr1)

	var arr3 []byte
	arr3 = append(arr3,[]byte("hello")...)
	fmt.Printf("%s\n",arr3)
	//变长的函数参数
	f1(0,1,2)
	//0 0
	//0 1
	//0 2
	//变长的函数参数
	b := []int{0,1,2}
	f2(b)
	//0 0
	//0 1
	//0 2
var req = RegisterReq{
	UserName:"xcx",
	PasswordNew:"1121",
	PasswordRepeat:"dasds",
	Email:"xasaaqqw.com",
}

err := validates(req)
fmt.Println(err)

}
