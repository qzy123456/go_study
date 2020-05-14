package main
import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(3 * time.Minute))
	request := make([]byte,1024)
	defer conn.Close()

	for {
		recv_len,err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if recv_len == 0 {
			break
		}
		recvData := strings.TrimSpace(string(request[:recv_len]))
		fmt.Println("recv_len : ",recv_len)
		fmt.Println("recv_data : " + recvData)
		daytime := time.Now().String()
		conn.Write([]byte(daytime + "\n"))
		request = make([]byte,1024)
	}
}

func main() {
	bindInfo := ":12345"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",bindInfo)
	checkError(err)
	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for {
		cc,err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(cc)
	}
}