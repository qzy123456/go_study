package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}
//定义服务
type SampleService struct {
}

//服务里的方法
func (this *SampleService) GetUserInfo(uid int64) error {
	fmt.Println(uid)
	return nil
}

func main() {
	//tcp,推荐
	server := rpc.NewTCPServer("tcp4://127.0.0.1:8082/")
	//注册func
	server.AddFunction("hello1", hello)
	//注册struct，命名空间是Sample
	server.AddInstanceMethods(&SampleService{}, rpc.Options{NameSpace: "Sample"})
	err := server.Start()
	if err != nil {
		fmt.Printf("start server fail, err:%v\n", err)
		return
	}
}