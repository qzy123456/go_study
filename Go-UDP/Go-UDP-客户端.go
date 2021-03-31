package main


import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 连接服务器
	conn,err := net.DialUDP("udp",nil,&net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 9090,
	})

	if err != nil {
		log.Println("Connect to udp server failed,err:",err)
		return
	}

	// 循环读取消息
	for {
		select {
		case <-time.After(time.Second * 20):
			for i:=0;i<10;i++{
				// 发送数据
				_,err := conn.Write([]byte(fmt.Sprintf("udp testing:%v",i)))
				if err != nil {
					fmt.Printf("Send data failed,err:",err)
					return
				}

				//接收数据
				result := make([]byte,1024)
				n,remoteAddr,err := conn.ReadFromUDP(result)
				if err != nil{
					fmt.Printf("Read from udp server failed ,err:",err)
					return
				}
				fmt.Printf("Recived msg from %s, data:%s \n",remoteAddr,string(result[:n]))
			}
		}
	}


}