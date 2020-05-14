package main

import "fmt"
//实名结构体
type People struct {
	name  string
	child *People
}
type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}
//打印消息类型，匿名结构体
func printtest(msg *struct{
	id int
	data string
})  {
	fmt.Printf("%T\n",msg)
}
func main()  {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	addr :=Address{
		"xs",
		"xsx",
		211,
		"weq",
	}
	fmt.Println(relation,"\n",addr)
	//实例化调用匿名结构体
	msg := &struct{
		id int
		data string
	}{
		1024,
		"data",
	}

	printtest(msg)
	//*struct { id int; data string }

}
