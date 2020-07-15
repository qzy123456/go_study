package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"math/rand"
	"net"
	"time"
)

var serverList []string
func GetConnect() (conn *zk.Conn, err error) {
	hosts := []string{"localhost:2181"}
	conn, _, err = zk.Connect(hosts, 5*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}


func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	fmt.Println("list:", list)
	return
}

//watch机制，服务器有断开或者重连，收到消息
func watchServerList(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errs := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				errs <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errs <- evt.Err
				return
			}
		}
	}()

	return snapshots, errs
}

//watch机制，监听配置文件变化的过程
func watchGetDat(conn *zk.Conn, path string) (chan []byte, chan error) {
	snapshots := make(chan []byte)
	errors := make(chan error)

	go func() {
		for {
			dataBuf, _, events, err := conn.GetW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- dataBuf
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
		}
	}()

	return snapshots, errors
}

func main() {
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err = GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//用来实时监听服务的上线与下线功能，serverList时刻保持最新的在线服务
	snapshots, errors := watchServerList(conn, "/go_servers")
	go func() {
		for {
			select {
			case serverList = <-snapshots:
				fmt.Printf("1111:%+v\n", serverList)
			case err := <-errors:
				fmt.Printf("2222:%+v\n", err)
			}
		}
	}()

	configs, errors := watchGetDat(conn, "/config")
	go func() {
		for {
			select {
			case configData := <-configs:
				fmt.Printf("333:%+v\n", string(configData))
			case err := <-errors:
				fmt.Printf("4444:%+v\n", err)
			}
		}
	}()
	for i := 0; i < 100; i++ {
		startClient()

		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	// service := "127.0.0.1:8899"
	//获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	//serverHost := "127.0.0.1:8899"
	fmt.Println("connect host: " + serverHost)
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	//checkError(err)
	conn, err := net.Dial("tcp", serverHost)
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	defer conn.Close()
	fmt.Println("connect ok")
	_, err = conn.Write([]byte("timestamp"))
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	fmt.Println("write ok")
	// result, err := ioutil.ReadAll(conn)
	// checkError(err)
	// fmt.Println("recv:", string(result))

	return
}

func getServerHost() (host string, err error) {
	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}
