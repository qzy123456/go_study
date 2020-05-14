package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type HelloService struct {}

func (p *HelloService) Hello(request string,reply *string)error  {
	*reply = "hello" + request
	return nil
}
func main()  {
	rpc.RegisterName("HelloService",new(HelloService))

	listener,err := net.Listen("tcp",":1234")
	if err != nil{
		fmt.Println("listen Tcp error",err)
	}
	//生成随机数
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		fmt.Println(rand1.Intn(15))
	}
	for {
		conn, err := listener.Accept()
		fmt.Println(conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("listen error", err)
		}
		//rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}


}