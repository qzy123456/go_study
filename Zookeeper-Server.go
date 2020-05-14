package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"net"
	"os"
	"time"
)

func main() {
	go starServer("127.0.0.1:9897")
	go starServer("127.0.0.1:9898")
	go starServer("127.0.0.1:9899")

	a := make(chan bool, 1)
	<-a
}
func GetConnects() (conn *zk.Conn, err error) {
	hosts := []string{"localhost:2181"}
	conn, _, err = zk.Connect(hosts, 5*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	// _, err = conn.Create("/config", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	// fmt.Println("err:", err)
	// stat, err := conn.Set("/config", []byte("hello world"), -1)
	// fmt.Println("stat:", stat)
	// fmt.Println("err:", err)
	// buf, stat, err := conn.Get("/config")
	// fmt.Println("buf:", string(buf))
	// fmt.Println("stat:", stat)
	// fmt.Println("err:", err)
	return
}
func RegistServers(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}
func starServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	if err != nil {
		fmt.Println("err:", err)
		//panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("err:", err)
		//panic(err)
	}

	//注册zk节点q
	conn, err := GetConnects()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	err = RegistServers(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn, port)
	}

	fmt.Println("aaaaaa")
}

func handleCient(conn net.Conn, port string) {
	fmt.Println("new client:", conn.RemoteAddr())
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Receive data from client:", string(buf[:length]))
		// conn.Write([]byte("hello world"))
	}
}

