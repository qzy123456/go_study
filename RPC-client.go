package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	coon,err := net.Dial("tcp","localhost:1234")
	if err!= nil{
		fmt.Println(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(coon))
	var reply string
	err = client.Call("HelloService.Hello","word",&reply)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(reply)
}
